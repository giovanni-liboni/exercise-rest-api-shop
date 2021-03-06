package repositories

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type orderRepository struct {
	db *sqlx.DB
}

type OrderRepository interface {
	GetAllOrders(ctx context.Context) ([]*entities.Order, error)
	GetOrdersByUserID(ctx context.Context, userID int64) ([]*entities.Order, error)
	GetOrder(ctx context.Context, id int64) (*entities.Order, error)
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) error
	GetOrderItems(ctx context.Context, orderID int64) ([]*entities.Item, error)
	GetOrdersByUserIDAndStatus(ctx context.Context, userID int64, status string) ([]*entities.Order, error)
}

func InitOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{db}
}

func (o orderRepository) GetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := o.db.SelectContext(ctx, &orders, "CALL sp_GetOrders()")
	return orders, err
}

func (o orderRepository) GetOrdersByUserID(ctx context.Context, userID int64) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := o.db.SelectContext(ctx, &orders, "CALL sp_GetOrdersByUserID(?)", userID)
	return orders, err
}

func (o orderRepository) GetOrder(ctx context.Context, id int64) (*entities.Order, error) {
	var order entities.Order
	err := o.db.GetContext(ctx, &order, "CALL sp_GetOrder(?)", id)
	return &order, err
}

func (o orderRepository) CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	query := "INSERT INTO orders (user_id, payment_method, payment_id, total_price, status) " +
		"VALUES  (:user_id, :payment_method, :payment_id, :total_price, :status)"
	res, err := o.db.NamedExecContext(ctx, query, order)
	if err != nil {
		return nil, err
	}

	// Return the order and the error
	order.ID, _ = res.LastInsertId()
	// If there are items in the order, create the order items
	if len(order.Items) > 0 {
		for _, item := range order.Items {
			_, err = o.db.ExecContext(ctx, "CALL sp_CreateOrderItem(?, ?, ?)", order.ID, item.ID, item.Price)
		}
	}
	return order, err
}

func (o orderRepository) UpdateOrder(ctx context.Context, order *entities.Order) error {
	_, err := o.db.NamedExecContext(ctx, "CALL sp_UpdateOrder(:id, :user_id, :payment_method, :total_price, :status, :payment_id)", order)
	return err
}

func (o orderRepository) GetOrderItems(ctx context.Context, orderID int64) ([]*entities.Item, error) {
	var items []*entities.Item
	err := o.db.SelectContext(ctx, &items, "CALL sp_GetOrderItems(?)", orderID)
	return items, err
}

func (o orderRepository) GetOrdersByUserIDAndStatus(ctx context.Context, userID int64, status string) ([]*entities.Order, error) {
	var orders []*entities.Order
	log.Infoln("CALL sp_GetOrdersByUserIDAndStatus(?, ?)", userID, status)
	err := o.db.SelectContext(ctx, &orders, "CALL sp_GetOrdersByUserIDAndStatus(?, ?)", userID, status)
	// print the orders
	for _, order := range orders {
		log.Infoln("Order: ", order)
	}
	return orders, err
}
