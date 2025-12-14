package consumer

import (
	"time"

	"github.com/darialissi/msa_big_tech/notification/internal/app/modules/inbox"

	"github.com/IBM/sarama"
)

func toInboxMessage(eventID string, msg *sarama.ConsumerMessage) *inbox.Message {
	return &inbox.Message{
		ID:         eventID,
		Topic:      msg.Topic,
		Partition:  msg.Partition,
		Offset:     msg.Offset,
		Payload:    msg.Value,
		Status:     inbox.StatusReceived,
		ReceivedAt: time.Now(),
	}
}
