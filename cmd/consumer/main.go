package main

import (
	"fmt"
	"github/jailtonjunior94/go-kafka/business/environments"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	environments.NewConfig()

	consume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": environments.BootstrapServer,
		"group.id":          environments.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	consume.Subscribe(environments.Topic, nil)

	for {
		message, err := consume.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", message.TopicPartition, string(message.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, message)
		}
	}
}
