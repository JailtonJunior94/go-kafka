package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github/jailtonjunior94/go-kafka/business/dtos"
	"github/jailtonjunior94/go-kafka/business/environments"
)

type ISlackService interface {
	SendMessage(req *dtos.SlackRequest) error
}

type SlackService struct {
}

func NewSlackService() ISlackService {
	return &SlackService{}
}

func (s SlackService) SendMessage(req *dtos.SlackRequest) error {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := http.Post(environments.SlackBaseUrl, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
