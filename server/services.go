package server

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/services"
)

type Services struct {
	ItemService  services.ItemService
	OrderService services.OrderService
	UserService  services.UserService
	StatService  services.StatService
}

func InitServices(repos *Repositories, config *config.Config) *Services {
	return &Services{
		ItemService:  services.InitItemService(repos.ItemRepository, repos.OrderRepository),
		OrderService: services.InitOrderService(repos.OrderRepository, config.StripeAPIKey),
		UserService:  services.InitUserService(repos.UserRepository, repos.OrderRepository),
		StatService:  services.InitStatService(repos.UserRepository, repos.OrderRepository, repos.ItemRepository, repos.StatRepository),
	}
}
