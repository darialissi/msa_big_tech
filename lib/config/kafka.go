package config

import (
	"errors"
)

type KfEnv struct {
	brokers      string
	fr_req_topic string
}

func KfConfig(mode string) *KfEnv {
	if mode == "dev" {
		return &KfEnv{
			brokers:      getEnv("KAFKA_BROKERS_DEV", ""),
			fr_req_topic: getEnv("KAFKA_FR_REQ_EVENTS_TOPIC_NAME_DEV", ""),
		}
	}
	return &KfEnv{
		brokers:      getEnv("KAFKA_BROKERS", ""),
		fr_req_topic: getEnv("KAFKA_FR_EVENTS_TOPIC_NAME", ""),
	}
}

func (env *KfEnv) Validate() error {
	if env.brokers == "" {
		return errors.New("No defined KAFKA brokers")
	}

	if env.fr_req_topic == "" {
		return errors.New("No defined KAFKA fr_req_topic")
	}

	return nil
}

func (env *KfEnv) GetBrokers() string {
	return env.brokers
}

func (env *KfEnv) GetFrReqTopic() string {
	return env.fr_req_topic
}
