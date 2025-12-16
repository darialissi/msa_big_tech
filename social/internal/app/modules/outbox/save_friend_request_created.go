package outbox

import (
	"context"
	"time"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

func (p *Processor) SaveFriendRequestCreated(ctx context.Context, req *models.FriendRequest) error {
	api := "SaveFriendRequestCreated"

	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: Marshal: %w", api, err)
	}

	e := Event{
		ID:            uuid.New(),
		AggregateType: AggregateTypeFriendRequest,
		AggregateID:   string(req.ID),
		EventType:     EventTypeFriendRequestCreated,
		Payload:       payload,
		CreatedAt:     time.Now().UTC(),
	}

	return p.Repository.SaveEvent(ctx, &e)
}
