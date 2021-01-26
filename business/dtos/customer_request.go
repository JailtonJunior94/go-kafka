package dtos

import "errors"

type CustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (customer *CustomerRequest) IsValid() error {
	if customer.Name == "" {
		return errors.New("O Nome é obrigatório")
	}

	if customer.Email == "" {
		return errors.New("O E-mail é obrigatório")
	}

	return nil
}
