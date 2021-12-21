package server

import "github.com/giovanni-liboni/exercise-rest-api-shop/handlers"

type Handlers struct {
	ItemHandler handlers.ItemHandler
	OrderHandler handlers.OrderHandler
	UserHandler handlers.UserHandler
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		ItemHandler: handlers.InitItemHandler(services.ItemService),
		OrderHandler: handlers.InitOrderHandler(services.OrderService, services.UserService),
		UserHandler: handlers.InitUserHandler(services.UserService, services.OrderService),
	}
}
