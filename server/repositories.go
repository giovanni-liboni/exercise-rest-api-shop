package server

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
}
}
