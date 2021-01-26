package entities

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        int64        `db:"Id"`
	Name      string       `db:"Name"`
	Email     string       `db:"Email"`
	CreatedAt time.Time    `db:"CreatedAt"`
	UpdatedAt sql.NullTime `db:"UpdatedAt"`
	Active    bool         `db:"Active"`
}
