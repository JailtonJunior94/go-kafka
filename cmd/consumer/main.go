package main

import (
	"encoding/json"
	"fmt"
	"github/jailtonjunior94/go-kafka/business/environments"
	"github/jailtonjunior94/go-kafka/business/messages"
	"github/jailtonjunior94/go-kafka/business/services"

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

	slack := services.NewSlackService()
	notification := services.NewNotificationService(slack)

	fmt.Println("🚀 Consumer is running")

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

			if err = notification.SendNotification(&kafkaMessage); err != nil {
				log.Error(fmt.Sprintf("[SEND NOTIFICATION]: %v", err))
			}

			tp, err := consume.CommitMessage(message)
			if err != nil {
				log.Error(fmt.Sprintf("[COMMIT MESSAGE]: %v", err))
			}
			log.Info(fmt.Sprintf("[TOPIC PARTITION]: %v", tp))
		}
	}
}
