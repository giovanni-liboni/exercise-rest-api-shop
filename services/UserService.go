package services

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type userService struct {
	userRepository repositories.UserRepository
	orderRepository repositories.OrderRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByID(ctx context.Context, id int64) (*entities.User, error)
}

func InitUserService(userRepository repositories.UserRepository, orderRepository repositories.OrderRepository) UserService {
	return &userService{
		userRepository: userRepository,
		orderRepository: orderRepository,
	}
}

// Check if the username contains only alphanumeric characters and underscores
var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString

// CreateUser creates a new user in the database. It checks if the username is a valid alphanumeric string.
// It encrypts the password using bcrypt. It returns an error if the username is not valid or if the password is empty.
func (u userService) CreateUser(ctx context.Context, user *entities.User) error {

	// Check if the username is alphanumeric
	if !isStringAlphabetic(user.Username) && user.Username != "" {
		return entities.ErrInvalidUsername
	}

	// Check if the password is empty
	if len(user.Password) == 0 {
		return entities.ErrEmptyPassword
	}

	// Encrypt the password
	userPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(userPassword)

	return u.userRepository.CreateUser(ctx, user)
}

func (u userService) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	user, err := u.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}