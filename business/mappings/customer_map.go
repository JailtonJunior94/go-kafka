package mappings

import (
	"github/jailtonjunior94/go-kafka/business/dtos"
	"github/jailtonjunior94/go-kafka/business/entities"
	"time"
)

func ToResponse(c entities.Customer) dtos.CustomerResponse {
	return dtos.CustomerResponse{
		ID:     c.ID,
		Name:   c.Name,
		Email:  c.Email,
		Active: c.Active,
	}
}

func ToListResponse(c []entities.Customer) []dtos.CustomerResponse {
	var list []dtos.CustomerResponse

	for _, c := range c {
		customer := dtos.CustomerResponse{
			ID:     c.ID,
			Name:   c.Name,
			Email:  c.Email,
			Active: c.Active,
		}
		list = append(list, customer)
	}

	return list
}

func ToEntity(r dtos.CustomerRequest) entities.Customer {
	return entities.Customer{
		Name:      r.Name,
		Email:     r.Email,
		CreatedAt: time.Now(),
		Active:    true,
	}
}
