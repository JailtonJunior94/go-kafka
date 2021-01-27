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
		log.Error(err)
	}

	if err = consume.Subscribe(environments.Topic, nil); err != nil {
		log.Error(err)
	}

	for {
		message, err := consume.ReadMessage(-1)
		if err != nil {
			log.Error(err)
		}

		if message.Value != nil {
			var kafkaMessage messages.KafkaMessage
			if err := json.Unmarshal(message.Value, &kafkaMessage); err != nil {
				log.Error(err)
			}

			if kafkaMessage.Payload.After != nil && kafkaMessage.Payload.Before == nil {
				log.Info(fmt.Sprintf("[INSERT] - %v\n", &kafkaMessage.Payload))
			} else {
				log.Info(fmt.Sprintf("[UPDATE] - %v\n", &kafkaMessage.Payload))
			}

			_, err = consume.CommitMessage(message)
			if err != nil {
				fmt.Printf("Erro ao comitar mensagem: %s", err)
			}
		}

		_, err = consume.CommitMessage(message)
		if err != nil {
			fmt.Printf("Erro ao comitar mensagem: %s", err)
		}
	}
}
