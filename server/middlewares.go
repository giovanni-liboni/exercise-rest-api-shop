package server

import "github.com/giovanni-liboni/exercise-rest-api-shop/config"

type Middlewares struct {
}

func InitMiddlewares(config *config.Config, repos *Repositories) *Middlewares {
	return &Middlewares{
	}
}
