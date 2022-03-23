package main

import (
	"encoding/json"

	"github/jailtonjunior94/go-kafka/business/environments"
	"github/jailtonjunior94/go-kafka/business/kafka"
	"github/jailtonjunior94/go-kafka/business/messages"
	"github/jailtonjunior94/go-kafka/business/services"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

func main() {
	environments.NewConfig()
	log.SetFormatter(&log.JSONFormatter{})

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": environments.BootstrapServer,
		"group.id":          environments.GroupId,
		"auto.offset.reset": "earliest",
	}
	topics := []string{environments.Topic}

	slack := services.NewSlackService()
	notification := services.NewNotificationService(slack)
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	log.Info("[🚀 Consumer is running]")

	for msg := range msgChan {
		var message *messages.KafkaMessage
		if err := json.Unmarshal(msg.Value, &message); err != nil {
			log.Error("[Não foi possível fazer o Unmarshal da mensagem]")
			continue
		}

		if err := notification.SendNotification(message); err != nil {
			log.Error("[Não foi possível enviar notificação para o Slack]")
			continue
		}
		log.Info("[Notificação enviada com sucesso!]")
	}
}
