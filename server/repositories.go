package server

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UserRepository repositories.UserRepository
	ItemRepository repositories.ItemRepository
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository: repositories.InitUserRepository(db),
		ItemRepository: repositories.InitItemRepository(db),
	}
}
