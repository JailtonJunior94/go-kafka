package repositories

import (
	"github/jailtonjunior94/go-kafka/business/entities"

	"github.com/jmoiron/sqlx"
)

type ICustomerReposity interface {
	Get() (customers []*entities.Customer, err error)
	GetById(id int64) (customer *entities.Customer, err error)
	Add(c *entities.Customer) (customer *entities.Customer, err error)
	Update(c *entities.Customer) (customer *entities.Customer, err error)
	Delete(id int64) error
}

type CustomerRepository struct {
	Db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) ICustomerReposity {
	return &CustomerRepository{Db: db}
}

func (r *CustomerRepository) Get() (customers []*entities.Customer, err error) {

	return nil, nil
}

func (r *CustomerRepository) GetById(id int64) (customer *entities.Customer, err error) {

	return nil, nil
}

func (r *CustomerRepository) Add(c *entities.Customer) (customer *entities.Customer, err error) {

	return nil, nil
}

func (r *CustomerRepository) Update(c *entities.Customer) (customer *entities.Customer, err error) {

	return nil, nil
}

func (r *CustomerRepository) Delete(id int64) error {

	return nil
}
