package server

import "github.com/giovanni-liboni/exercise-rest-api-shop/services"

type Services struct {
	ItemService services.ItemService
}

func InitServices(repos *Repositories) *Services {
	return &Services{
		ItemService: services.InitItemService(repos.ItemRepository),
	}
}
