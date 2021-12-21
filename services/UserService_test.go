package services

import (
	"context"
	"database/sql"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initUserService(db *sql.DB) UserService {
	// Create a new user repository
	userRepository := repositories.InitUserRepository(sqlx.NewDb(db, "mysql"))

	// Initialize the order repository
	orderRepository := repositories.InitOrderRepository(sqlx.NewDb(db, "mysql"))

	// Initialize the user service
	return InitUserService(userRepository, orderRepository)
}

func TestUserService_CreateUser(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	// Initialize the user service
	userService := initUserService(db)

	// Create a new user
	user := entities.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Password:  "password",
		Username: "johndoe",
	}

	// Create the user
	err = userService.CreateUser(context.TODO(), &user)
	if err != nil {
		t.Fatalf("mysql: failed to create a user, err: %s", err)
	}
}

func TestUserService_CreateUser_UserAlreadyExists(t *testing.T) {
	db, err := sql.Open("mysql_txdb", "testDataSource")
	if err != nil {
		t.Fatalf("mysql: failed to open a mysql connection, have you run 'make test'? err: %s", err)
	}
	defer db.Close()

	// Initialize the user service
	userService := initUserService(db)

	// Create a new user
	user := entities.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "test.doe@gmail.com",
		Password:  "password",
		Username: "test",
	}

	// Create the user
	err = userService.CreateUser(context.TODO(), &user)

	assert.ErrorIs(t, err, entities.ErrUserAlreadyExists)
}

