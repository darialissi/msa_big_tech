package outbox

import (
	"context"

	uc "github.com/darialissi/msa_big_tech/social/internal/app/usecases"
)

// Проверка удовлетворению интерфейсу orders.OutboxRepository
var _ uc.OutboxRepository = (*Processor)(nil)

type (
	// Repository - репозиторий outbox
	Repository interface {
		SaveEvent(ctx context.Context, e *Event) error
		SearchEvents(ctx context.Context, opts ...SearchEventsOption) []*Event
		UpdateEvents(ctx context.Context, opts ...UpdateEventsOption) error
	}

	// TransactionManager
	TransactionManager interface {
		RunRepeatableRead(ctx context.Context, f func(tctx context.Context) error) error
	}
)

// Deps - зависимости
type Deps struct {
	Repository Repository
}

// Processor - ...
type Processor struct {
	Deps
}

// NewProcessor - ...
func NewProcessor(d Deps) *Processor {
	return &Processor{
		Deps: d,
	}
}
