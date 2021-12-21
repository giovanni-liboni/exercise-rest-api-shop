package repositories

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_GetAllUsers(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	timestamp := time.Now()
	// Define the user rows to be returned by the database
	rows := sqlmock.NewRows([]string{"firstname", "lastname", "email", "username", "password", "role", "created_at", "updated_at"}).
		AddRow("John", "Doe", "john.doe@gmail.com", "johndoe", "password", "user", timestamp, timestamp).
		AddRow("Jane", "Doe", "jane.doe@gmail.com", "janedoe", "password", "user", timestamp, timestamp).
		AddRow("John", "Smith", "john.smith@gmail.com", "johnsmith", "password", "user", timestamp, timestamp)

	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetUsers").WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitUserRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetAllItems method
	items, err := repo.GetAllUsers(context.TODO())
	if err != nil {
		t.Errorf("error was not expected while getting users: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Assert that the returned items are not empty
	assert.NotEmpty(t, items)
}

func TestUserRepository_GetUserByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	timestamp := time.Now()
	// Define the user rows to be returned by the database
	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "username", "password", "role", "created_at", "updated_at"}).
		AddRow(1, "John", "Doe", "john.doe@gmail.com", "johndoe", "password", "user", timestamp, timestamp)

	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetUserByID").WithArgs(1).WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitUserRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetUserByID method
	item, err := repo.GetUserByID(context.TODO(), 1)
	if err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check assertions
	assert.NotEmpty(t, item)
	assert.Equal(t, int64(1), item.ID)
	assert.Equal(t, "John", item.Firstname)
	assert.Equal(t, "Doe", item.Lastname)
	assert.Equal(t, "johndoe", item.Username)

}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	timestamp := time.Now()
	// Define the user rows to be returned by the database
	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "username", "password", "role", "created_at", "updated_at"}).
		AddRow(1, "John", "Doe", "john.doe@gmail.com", "johndoe", "password", "user", timestamp, timestamp)

	// Expect the query to be called with the correct SQL
	mock.ExpectQuery("CALL sp_GetUserByUsername").WithArgs("johndoe").WillReturnRows(rows)

	// Create a new ItemRepository with the database connection
	repo := InitUserRepository(sqlx.NewDb(db, "mysql"))

	// Call the GetUserByUsername method
	item, err := repo.GetUserByUsername(context.TODO(), "johndoe")
	if err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	// Check that the expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check assertions
	assert.NotEmpty(t, item)
	assert.Equal(t, int64(1), item.ID)
	assert.Equal(t, "John", item.Firstname)
	assert.Equal(t, "Doe", item.Lastname)
	assert.Equal(t, "johndoe", item.Username)
}

func TestUserRepository_CreateUser_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the user
	user := &entities.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Username:  "johndoe",
		Password:  "password",
		Role:      "user",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("CALL sp_GetUserByUsername(?)").WithArgs("johndoe").WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("CALL sp_CreateUser").WithArgs("John", "Doe", "john.doe@gmail.com", "johndoe", "password").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepository := InitUserRepository(sqlx.NewDb(db, "mysql"))
	err = userRepository.CreateUser(context.TODO(), user)
	if err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_CreateUser_Found(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the user
	user := &entities.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Username:  "johndoe",
		Password:  "password",
		Role:      "user",
	}

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "username", "password", "role", "created_at", "updated_at"}).
		AddRow(1, "John", "Doe", "john.doe@gmail.com", "johndoe", "password", "user", time.Now(), time.Now())

	mock.ExpectBegin()
	mock.ExpectQuery("CALL sp_GetUserByUsername(?)").WithArgs("johndoe").WillReturnRows(rows)
	mock.ExpectRollback()

	userRepository := InitUserRepository(sqlx.NewDb(db, "mysql"))
	err = userRepository.CreateUser(context.TODO(), user)
	if err != entities.ErrUserAlreadyExists {
		t.Errorf("error was not expected while creating user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.ErrorIs(t, entities.ErrUserAlreadyExists, err)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the user
	user := &entities.User{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Username:  "johndoe",
		Password:  "password",
		Role:      "user",
	}

	mock.ExpectExec("CALL sp_UpdateUser").WithArgs(1, "John", "Doe", "john.doe@gmail.com", "johndoe", "password", "user").WillReturnResult(sqlmock.NewResult(1, 1))

	userRepository := InitUserRepository(sqlx.NewDb(db, "mysql"))
	err = userRepository.UpdateUser(context.TODO(), user)

	if err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_DeleteUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CALL sp_DeleteUser").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	userRepository := InitUserRepository(sqlx.NewDb(db, "mysql"))
	err = userRepository.DeleteUser(context.TODO(), 1)

	if err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
