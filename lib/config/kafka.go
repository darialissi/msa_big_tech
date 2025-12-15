package config

import (
	"errors"
)

type consumer struct {
	group string
	name string
}

type KfEnv struct {
	brokers      string
	fr_req_topic string
	consumer
}

func KfConfig(mode string) *KfEnv {
	if mode == "dev" {
		return &KfEnv{
			brokers:      getEnv("KAFKA_BROKERS_DEV", ""),
			fr_req_topic: getEnv("KAFKA_FR_REQ_EVENTS_TOPIC_NAME_DEV", ""),
			consumer: consumer{
				group: getEnv("KAFKA_CONSUMER_GROUP_DEV", ""),
				name: getEnv("KAFKA_CONSUMER_NAME_DEV", ""),
			},
		}
	}
	return &KfEnv{
		brokers:      getEnv("KAFKA_BROKERS", ""),
		fr_req_topic: getEnv("KAFKA_FR_REQ_EVENTS_TOPIC_NAME", ""),
		consumer: consumer{
			group: getEnv("KAFKA_CONSUMER_GROUP", ""),
			name: getEnv("KAFKA_CONSUMER_NAME", ""),
		},
	}
}

func (env *KfEnv) Validate() error {
	if env.brokers == "" {
		return errors.New("No defined KAFKA brokers")
	}

	return nil
}

func (env *KfEnv) GetBrokers() string {
	return env.brokers
}

func (env *KfEnv) GetFrReqTopic() string {
	return env.fr_req_topic
}

func (env *KfEnv) GetConsumerGroup() string {
	return env.consumer.group
}

func (env *KfEnv) GetConsumerName() string {
	return env.consumer.name
}
