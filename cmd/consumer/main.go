package main

import (
	"encoding/json"
	"fmt"
	"github/jailtonjunior94/go-kafka/business/environments"
	"github/jailtonjunior94/go-kafka/business/messages"

	log "github.com/sirupsen/logrus"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	environments.NewConfig()
	log.SetFormatter(&log.JSONFormatter{})

	consume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": environments.BootstrapServer,
		"group.id":          environments.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Error(fmt.Sprintf("[NEW CONSUMER]: %v", err))
	}

	if err = consume.Subscribe(environments.Topic, nil); err != nil {
		log.Error(fmt.Sprintf("[SUBSCRIBE]: %v", err))
	}

	for {
		message, err := consume.ReadMessage(-1)
		if err != nil {
			log.Error(fmt.Sprintf("[READ MESSAGE]: %v", err))
		}

		if message.Value != nil {
			var kafkaMessage messages.KafkaMessage
			if err := json.Unmarshal(message.Value, &kafkaMessage); err != nil {
				log.Error(fmt.Sprintf("[JSON UNMARSHAL]: %v", err))
			}

			if kafkaMessage.Payload.After != nil && kafkaMessage.Payload.Before == nil {
				log.Info(fmt.Sprintf("[INSERT]: %v", *kafkaMessage.Payload.After))
			}

			if kafkaMessage.Payload.After != nil && kafkaMessage.Payload.Before != nil {
				log.Info(fmt.Sprintf("[UPDATE]: %v", *kafkaMessage.Payload.After))
			}

			if kafkaMessage.Payload.After == nil && kafkaMessage.Payload.Before != nil {
				log.Info(fmt.Sprintf("[DELETE]: %v", *kafkaMessage.Payload.Before))
			}

			tp, err := consume.CommitMessage(message)
			if err != nil {
				log.Error(fmt.Sprintf("[COMMIT MESSAGE]: %v", err))
			}
			log.Info(fmt.Sprintf("[TOPIC PARTITION]: %v", tp))
		}
	}
}
