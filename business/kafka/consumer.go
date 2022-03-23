package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan<- *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		log.Error(err)
		return err
	}

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
		log.Error(err)
	}
}
