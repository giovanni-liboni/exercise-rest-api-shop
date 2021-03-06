package server

import (
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/middlewares"
)

type Middlewares struct {
	AuthMiddleware      *middlewares.AuthMiddleware
	GroupAuthMiddleware *middlewares.GroupAuthMiddleware
}

func InitMiddlewares(config *config.Config, repos *Repositories) *Middlewares {
	return &Middlewares{
		AuthMiddleware:      middlewares.InitAuthMiddleware(config, repos.UserRepository),
		GroupAuthMiddleware: middlewares.InitGroupAuthMiddleware(repos.UserRepository),
	}
}
