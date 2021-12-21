package repositories

import (
	"context"
	"database/sql"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*entities.User, error)
	GetUserByID(ctx context.Context, id int64) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id int64) error
}

func InitUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	var users []*entities.User
	err := u.db.SelectContext(ctx, &users, "CALL sp_GetUsers()")
	return users, err
}

func (u userRepository) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	var user entities.User
	err := u.db.GetContext(ctx, &user, "CALL sp_GetUserByID(?)", id)
	return &user, err
}

func (u userRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := u.db.GetContext(ctx, &user, "CALL sp_GetUserByUsername(?)", username)
	return &user, err
}

func (u userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var userExist entities.User
	err = tx.GetContext(ctx, &userExist, "CALL sp_GetUserByUsername(?)", user.Username)
	if err == sql.ErrNoRows {
		_, err = tx.NamedExecContext(ctx, "CALL sp_CreateUser(:firstname, :lastname, :email, :username, :password)", user)
		if err != nil {
			return err
		}
		return tx.Commit()
	} else if err != nil {
		return err
	}
	// Return error if user already exists
	return entities.ErrUserAlreadyExists
}

func (u userRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	_, err := u.db.NamedExecContext(ctx, "CALL sp_UpdateUser(:id, :firstname, :lastname, :email, :username, :password, :role)", user)
	return err
}

func (u userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := u.db.ExecContext(ctx, "CALL sp_DeleteUser(?)", id)
	return err
}
