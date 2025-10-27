package consumer

import (
	"time"

	"github.com/IBM/sarama"
)


type InboxConsumer struct {
	group        sarama.ConsumerGroup
	dedup        Deduper
	handler      Handler
	batchSize    int
	batchTimeout time.Duration
	consumerName string
}