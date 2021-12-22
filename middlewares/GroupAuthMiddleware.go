package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"net/http"
)

type GroupAuthMiddleware struct {
	userRepository repositories.UserRepository
}

func InitGroupAuthMiddleware(userRepository repositories.UserRepository) *GroupAuthMiddleware {
	return &GroupAuthMiddleware{
		userRepository: userRepository,
	}
}

func (gam *GroupAuthMiddleware) MiddlewareFunc(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userID")
		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userEntity, err := gam.userRepository.GetUserByID(c, user.(*entities.User).ID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		switch role {
		case "user":
			if userEntity.Role != "user" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
			return
		case "admin":
			if userEntity.Role != "admin" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
			return
		default:
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
