package outbox

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

func (p *Processor) SaveFriendRequestCreatedID(ctx context.Context, id models.FriendRequestID) error {
	e := Event{
		ID:            uuid.New(),
		AggregateType: AggregateTypeFriendRequest,
		AggregateID:   string(id),
		EventType:     EventTypeFriendRequestCreated,
		Payload:       nil,
		CreatedAt:     time.Now().UTC(),
	}

	return p.Repository.SaveEvent(ctx, &e)
}
