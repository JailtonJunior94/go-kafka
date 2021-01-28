package services

import (
	"fmt"
	"github/jailtonjunior94/go-kafka/business/dtos"
	"github/jailtonjunior94/go-kafka/business/messages"

	log "github.com/sirupsen/logrus"
)

type INotificationService interface {
	SendNotification(m *messages.KafkaMessage) error
}

type NotificationService struct {
	SlackService ISlackService
}

func NewNotificationService(slackService ISlackService) INotificationService {
	return &NotificationService{SlackService: slackService}
}

func (n NotificationService) SendNotification(m *messages.KafkaMessage) error {
	if m.Payload.After != nil && m.Payload.Before == nil {
		log.Info(fmt.Sprintf("[INSERT]: %v", *m.Payload.After))

		text := fmt.Sprintf(`[INSERT] - Cliente: 
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.Payload.After.Id, m.Payload.After.Name, m.Payload.After.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	if m.Payload.After != nil && m.Payload.Before != nil {
		log.Info(fmt.Sprintf("[UPDATE]: %v", *m.Payload.After))

		text := fmt.Sprintf(`[UPDATE] - Cliente: 
		      [BEFORE]
				ID: %v
				Nome: %s
				E-mail: %s
			   [AFTER]
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.Payload.Before.Id, m.Payload.Before.Name, m.Payload.Before.Email,
			m.Payload.After.Id, m.Payload.After.Name, m.Payload.After.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	if m.Payload.After == nil && m.Payload.Before != nil {
		log.Info(fmt.Sprintf("[DELETE]: %v", *m.Payload.Before))

		text := fmt.Sprintf(`[DELETE] - Cliente: 
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.Payload.Before.Id, m.Payload.Before.Name, m.Payload.Before.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	return nil
}
