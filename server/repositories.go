package server

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UserRepository  repositories.UserRepository
	ItemRepository  repositories.ItemRepository
	OrderRepository repositories.OrderRepository
	StatRepository  repositories.StatRepository
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:  repositories.InitUserRepository(db),
		ItemRepository:  repositories.InitItemRepository(db),
		OrderRepository: repositories.InitOrderRepository(db),
		StatRepository:  repositories.InitStatRepository(db),
	}
}
