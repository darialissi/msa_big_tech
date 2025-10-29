package outbox

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/modules/outbox"
)

type outboxEventRow struct {
	ID            uuid.UUID           `db:"id"`
	AggregateType string              `db:"aggregate_type"`
	AggregateID   string              `db:"aggregate_id"`
	EventType     string              `db:"event_type"`
	Payload       []byte              `db:"payload"` // JSONB
	CreatedAt     time.Time           `db:"created_at"`
	PublishedAt   sql.Null[time.Time] `db:"published_at"`
	RetryCount    int                 `db:"retry_count"`
	NextAttemptAt sql.Null[time.Time] `db:"next_attempt_at"`
}

func (e *outboxEventRow) mapFields() map[string]any {
	return map[string]any{
		outboxTableColumnID:            e.ID,
		outboxTableColumnAggType:       e.AggregateType,
		outboxTableColumnAggID:         e.AggregateID,
		outboxTableColumnEventType:     e.EventType,
		outboxTableColumnPayload:       e.Payload,
		outboxTableColumnCreatedAt:     e.CreatedAt,
		outboxTableColumnPublishedAt:   e.PublishedAt,
		outboxTableColumnRetryCount:    e.RetryCount,
		outboxTableColumnNextAttemptAt: e.NextAttemptAt,
	}
}

func (e *outboxEventRow) Values(columns ...string) []any {
	m := e.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := m[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

// ToModel конвертирует outboxEventRow (sql.Null*) в доменную модель models.Event (*string/*time.Time).
func ToModel(r *outboxEventRow) *outbox.Event {
	if r == nil {
		return nil
	}

	var publishedAt *time.Time
	if r.PublishedAt.Valid {
		t := r.PublishedAt.V
		publishedAt = &t
	}

	return &outbox.Event{
		ID:            r.ID,
		AggregateType: outbox.AggregateType(r.AggregateType),
		AggregateID:   r.AggregateID,
		EventType:     outbox.EventType(r.EventType),
		Payload:       r.Payload,
		CreatedAt:     r.CreatedAt,
		PublishedAt:   publishedAt,
		RetryCount:    r.RetryCount,
	}
}

// FromModel конвертирует доменную модель в outboxEventRow для INSERT/UPDATE/DELETE
func FromModel(e *outbox.Event) (outboxEventRow, error) {
	if e == nil {
		return outboxEventRow{}, fmt.Errorf("model is nil")
	}

	return outboxEventRow{
		ID:            e.ID,
		AggregateType: string(e.AggregateType),
		AggregateID:   e.AggregateID,
		EventType:     string(e.EventType),
		Payload:       notnullJSON(e.Payload),
		CreatedAt:     e.CreatedAt,
		PublishedAt: func(t *time.Time) sql.Null[time.Time] {
			if e.PublishedAt != nil {
				return sql.Null[time.Time]{V: *e.PublishedAt, Valid: true}
			}
			return sql.Null[time.Time]{}
		}(&e.CreatedAt),
		RetryCount: e.RetryCount,
	}, nil
}

func notnullJSON(data []byte) []byte {
	if data == nil {
		return []byte("[]")
	}
	return data
}
