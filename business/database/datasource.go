package database

import (
	"github/jailtonjunior94/go-kafka/business/environments"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type SqlConnection struct {
	Db *sqlx.DB
}

func NewConnection() (SqlConnection, error) {
	db, err := sqlx.Connect("sqlserver", environments.ConnectionString)
	if err != nil {
		return SqlConnection{}, err
	}

	if err = db.Ping(); err != nil {
		return SqlConnection{}, err
	}

	return SqlConnection{Db: db}, nil
}
