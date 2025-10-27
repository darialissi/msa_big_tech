package consumer

import (
	"context"

	"github.com/IBM/sarama"
)

type Handler interface {
	Handle(ctx context.Context, msg *sarama.ConsumerMessage) error
}