package outbox

import (
	"context"
	"fmt"
	"time"

	"github.com/darialissi/msa_big_tech/lib/postgres"
	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

func (r *Repository) SaveFriendRequestCreatedID(ctx context.Context, id models.FriendRequestID) error {

	e := models.Event{
		ID:            uuid.New(),
		AggregateType: models.AggregateTypeFriendRequest,
		AggregateID:   string(id),
		EventType:     models.EventTypeFriendRequestCreated,
		Payload:       nil,
		CreatedAt:     time.Now().UTC(),
	}

	return r.saveEvent(ctx, &e)
}

func (r *Repository) SaveFriendRequestUpdatedID(ctx context.Context, id models.FriendRequestID) error {

	e := models.Event{
		ID:            uuid.New(),
		AggregateType: models.AggregateTypeFriendRequest,
		AggregateID:   string(id),
		EventType:     models.EventTypeFriendRequestUpdated,
		Payload:       nil,
		CreatedAt:     time.Now().UTC(),
	}

	return r.saveEvent(ctx, &e)
}

func (r *Repository) saveEvent(ctx context.Context, in *models.Event) error {
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
