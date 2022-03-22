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
	if m.After != nil && m.Before == nil {
		log.Info(fmt.Sprintf("[INSERT]: %v", *m.After))

		text := fmt.Sprintf(`[INSERT] - Cliente: 
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.After.Id, m.After.Name, m.After.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	if m.After != nil && m.Before != nil {
		log.Info(fmt.Sprintf("[UPDATE]: %v", *m.After))

		text := fmt.Sprintf(`[UPDATE] - Cliente: 
		      [BEFORE]
				ID: %v
				Nome: %s
				E-mail: %s
			   [AFTER]
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.Before.Id, m.Before.Name, m.Before.Email,
			m.After.Id, m.After.Name, m.After.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	if m.After == nil && m.Before != nil {
		log.Info(fmt.Sprintf("[DELETE]: %v", *m.Before))

		text := fmt.Sprintf(`[DELETE] - Cliente: 
				ID: %v
				Nome: %s
				E-mail: %s
		`, m.Before.Id, m.Before.Name, m.Before.Email)

		if err := n.SlackService.SendMessage(&dtos.SlackRequest{Text: text}); err != nil {
			log.Error(fmt.Sprintf("[SLACK]: %v", err))
		}
	}

	return nil
}
