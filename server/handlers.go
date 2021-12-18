package server

import "github.com/giovanni-liboni/exercise-rest-api-shop/handlers"

type Handlers struct {
	ItemHandler handlers.ItemHandler
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		ItemHandler: handlers.InitItemHandler(services.ItemService),
	}
}
