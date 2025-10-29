package outbox

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/darialissi/msa_big_tech/lib/postgres"

	"github.com/darialissi/msa_big_tech/social/internal/app/modules/outbox"
)

func (r *Repository) SaveEvent(ctx context.Context, in *outbox.Event) error {
	const api = "outbox.Repository.SaveEvent"

	row, err := FromModel(in)

	if err != nil {
		return err
	}

	qb := r.sb.Insert(tableOutboxEvents).
		Columns(tableOutboxEventsColumns...).
		Values(row.Values(tableOutboxEventsColumns...)...)

	pool := r.db.GetQueryEngine(ctx)

	if _, err := pool.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}

// SearchEvents выбирает события из outbox по заданным опциям.
// Возвращает пустой срез при ошибке (сигнатура без error).
func (r *Repository) SearchEvents(ctx context.Context, opts ...outbox.SearchEventsOption) []*outbox.Event {
	o := outbox.CollectSearchEventsOptions(opts...)

	// Базовый селект
	qb := r.sb.
		Select(tableOutboxEventsColumns...).
		From(tableOutboxEvents).
		OrderBy(outboxTableColumnCreatedAt).
		Limit(uint64(o.Limit))

	// Фильтры
	if o.OnlyUnpublished {
		qb = qb.Where(squirrel.Eq{outboxTableColumnPublishedAt: nil}) // IS NULL
	}
	// retry_count <= MaxRetryCount
	qb = qb.Where(squirrel.LtOrEq{outboxTableColumnRetryCount: o.MaxRetryCount})

	if o.AggregateType != nil {
		qb = qb.Where(squirrel.Eq{outboxTableColumnAggType: string(*o.AggregateType)})
	}
	if o.EventType != nil {
		qb = qb.Where(squirrel.Eq{outboxTableColumnEventType: string(*o.EventType)})
	}
	if o.NotBefore != nil {
		qb = qb.Where(squirrel.GtOrEq{outboxTableColumnCreatedAt: *o.NotBefore})
	}
	if o.NotAfter != nil {
		qb = qb.Where(squirrel.LtOrEq{outboxTableColumnCreatedAt: *o.NotAfter})
	}
	if o.DueAt != nil {
		// next_attempt_at IS NULL OR next_attempt_at <= dueAt
		qb = qb.Where(
			squirrel.Or{
				squirrel.Eq{outboxTableColumnNextAttemptAt: nil},
				squirrel.LtOrEq{outboxTableColumnNextAttemptAt: *o.DueAt},
			},
		)
	}

	// Блокировка строк для конкурентных воркеров
	if o.WithLock {
		qb = qb.Suffix("FOR UPDATE SKIP LOCKED")
	}

	// Выполнение
	pool := r.db.GetQueryEngine(ctx)
	var rows []outboxEventRow
	if err := pool.Selectx(ctx, &rows, qb); err != nil {
		return nil
	}

	// Маппинг в доменную модель
	events := make([]*outbox.Event, 0, len(rows))

	for i, outRow := range rows {
		events[i] = ToModel(&outRow)
	}
	return events
}

func (r *Repository) UpdateEvents(ctx context.Context, opts ...outbox.UpdateEventsOption) error {
	const api = "outbox.Repository.UpdateEvents"

	o := outbox.CollectUpdateEventsOptions(opts...)

	// защита от noop
	if o.SetPublishedAt == nil && o.IncRetryBy == 0 && o.SetNextAttemptAt == nil {
		return nil
	}

	qb := r.sb.
		Update(tableOutboxEvents)

	// setters
	if o.SetPublishedAt != nil {
		qb = qb.Set(outboxTableColumnPublishedAt, *o.SetPublishedAt)
	}
	if o.IncRetryBy > 0 {
		qb = qb.Set(outboxTableColumnRetryCount, squirrel.Expr(outboxTableColumnRetryCount+" + ?", o.IncRetryBy))
	}
	if o.SetNextAttemptAt != nil {
		qb = qb.Set(outboxTableColumnNextAttemptAt, *o.SetNextAttemptAt)
	}

	// filters (для partition pruning)
	if o.AggregateType != nil {
		qb = qb.Where(squirrel.Eq{outboxTableColumnAggType: string(*o.AggregateType)})
	}
	if len(o.IDs) > 0 {
		qb = qb.Where(squirrel.Eq{outboxTableColumnID: o.IDs}) // id IN (...)
	}
	if o.NotBefore != nil {
		qb = qb.Where(squirrel.GtOrEq{outboxTableColumnCreatedAt: *o.NotBefore})
	}
	if o.NotAfter != nil {
		qb = qb.Where(squirrel.LtOrEq{outboxTableColumnCreatedAt: *o.NotAfter})
	}
	if o.OnlyUnpublished {
		qb = qb.Where(squirrel.Eq{outboxTableColumnPublishedAt: nil})
	}

	pool := r.db.GetQueryEngine(ctx)
	if _, err := pool.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}
	return nil
}
