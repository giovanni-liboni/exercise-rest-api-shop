package middlewares

import (
	"database/sql"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/giovanni-liboni/exercise-rest-api-shop/config"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AuthMiddleware struct {
	userRepository repositories.UserRepository
	Middleware     *jwt.GinJWTMiddleware
}

func InitAuthMiddleware(config *config.Config, userRepository repositories.UserRepository) *AuthMiddleware {
	var (
		identityKey = "username"
	)
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "rest-api",
		Key:         []byte(config.Jwt.Secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entities.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &entities.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Set a "login" flag in the Gin context to indicate that this is a login
			// request. This flag is needed in the unauthorized handler.
			c.Set("login", true)
			// Parse request data.
			var req login
			if err := c.BindJSON(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			// Get user profile from database.
			user, err := userRepository.GetUserByUsername(c, req.Username)
			if err == sql.ErrNoRows {
				return nil, jwt.ErrFailedAuthentication
			}
			if err != nil {
				// some internal error occurred.
				return nil, err
			}

			// Check if password is correct.
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
			if err == bcrypt.ErrMismatchedHashAndPassword {
				return nil, jwt.ErrFailedAuthentication
			}
			if err != nil {
				// some internal error occurred.
				return nil, err
			}
			return user, nil
		},
		// Unauthorized is called when a request to any endpoint is not
		// authorized.
		Unauthorized: func(c *gin.Context, code int, message string) {
			// Check if this is a login request. If it is, set the response status
			// code based on the error message.
			_, isLoginRequest := c.Get("login")
			if isLoginRequest {
				switch message {
				case jwt.ErrMissingLoginValues.Error():
					code = http.StatusBadRequest
				case jwt.ErrFailedAuthentication.Error():
					code = http.StatusUnauthorized
				default:
					code = http.StatusInternalServerError
				}
			}
			c.AbortWithError(code, errors.New(message))
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return &AuthMiddleware{
		Middleware:     authMiddleware,
		userRepository: userRepository,
	}
}
