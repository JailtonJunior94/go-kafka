package repositories

import (
	"database/sql"

	"github/jailtonjunior94/go-kafka/business/entities"

	"github.com/jmoiron/sqlx"
)

type ICustomerReposity interface {
	Get() (customers []entities.Customer, err error)
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

func (r CustomerRepository) Get() (customers []entities.Customer, err error) {
	query := `SELECT
				Id,
				Name,
				Email,
				CreatedAt,
				UpdatedAt,
				Active
			FROM
				dbo.Customers (NOLOCK)`

	if err := r.Db.Select(&customers, query); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r CustomerRepository) GetById(id int64) (customer *entities.Customer, err error) {
	query := `SELECT
				Id,
				Name,
				Email,
				CreatedAt,
				UpdatedAt,
				Active
			FROM
				dbo.Customers (NOLOCK)
			WHERE
				Id = @id`

	var customers []entities.Customer
	if err := r.Db.Select(&customers, query, sql.Named("id", id)); err != nil {
		return nil, err
	}

	if len(customers) == 0 {
		return nil, nil
	}

	return &customers[0], nil
}

func (r CustomerRepository) Add(c *entities.Customer) (customer *entities.Customer, err error) {
	query := `INSERT INTO
				dbo.Customers
			VALUES
				(@name, @email, @createdAt, @updatedAt, @active); SELECT SCOPE_IDENTITY()`

	s, err := r.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	lastInsertId := 0
	if err = s.QueryRow(sql.Named("name", c.Name),
		sql.Named("email", c.Email),
		sql.Named("createdAt", c.CreatedAt),
		sql.Named("updatedAt", c.UpdatedAt),
		sql.Named("active", c.Active)).Scan(&lastInsertId); err != nil {
		return nil, err
	}

	return &entities.Customer{
		ID:        int64(lastInsertId),
		Name:      c.Name,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Active:    c.Active,
	}, nil
}

func (r CustomerRepository) Update(c *entities.Customer) (customer *entities.Customer, err error) {
	query := `UPDATE
				dbo.Customers
			SET
				Name = @name,
				Email = @email,
				UpdatedAt = @updatedAt
			WHERE
				Id = @id`

	s, err := r.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	result, err := s.Exec(sql.Named("name", c.Name),
		sql.Named("email", c.Email),
		sql.Named("updatedAt", c.UpdatedAt.Time),
		sql.Named("id", c.ID))

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return nil, err
	}

	return c, nil
}

func (r CustomerRepository) Delete(id int64) error {
	query := `DELETE FROM
				dbo.Customers
			WHERE
				Id = @id`

	s, err := r.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer s.Close()

	result, err := s.Exec(sql.Named("id", id))
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return err
	}

	return nil
}
