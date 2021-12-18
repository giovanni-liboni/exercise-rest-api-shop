package repositories

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db *sqlx.DB
}

type OrderRepository interface {
	GetAllOrders(ctx context.Context) ([]*entities.Order, error)
	GetOrder(ctx context.Context, id int) (*entities.Order, error)
	CreateOrder(ctx context.Context, order *entities.Order) error
	UpdateOrder(ctx context.Context, order *entities.Order) error
}

func InitOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{db}
}

func (o orderRepository) GetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := o.db.SelectContext(ctx, &orders, "CALL sp_GetOrders()")
	return orders, err
}

func (o orderRepository) GetOrder(ctx context.Context, id int) (*entities.Order, error) {
	var order entities.Order
	err := o.db.GetContext(ctx, &order, "CALL sp_GetOrder(?)", id)
	return &order, err
}

func (o orderRepository) CreateOrder(ctx context.Context, order *entities.Order) error {
	_, err := o.db.NamedExecContext(ctx, "CALL sp_CreateOrder(:userID, :payment_method, :paymentID, :total_price, :status)", order)
	return err
}

func (o orderRepository) UpdateOrder(ctx context.Context, order *entities.Order) error {
	_, err := o.db.NamedExecContext(ctx, "CALL sp_UpdateOrder(:id, :userID, :payment_method, :total_price, :status, :paymentID)", order)
	return err
}
