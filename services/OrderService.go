package services

import (
	"context"
	"github.com/giovanni-liboni/exercise-rest-api-shop/entities"
	"github.com/giovanni-liboni/exercise-rest-api-shop/repositories"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

type orderService struct {
	orderRepository repositories.OrderRepository
	stripeAPIKey   string
}

type OrderService interface {
	GetOrder(ctx context.Context, id int64) (*entities.Order, error)
	GetOrders(ctx context.Context) ([]*entities.Order, error)
	GetOrdersByUser(ctx context.Context, userID int64) ([]*entities.Order, error)
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	PayOrder(ctx context.Context, id int64, user *entities.User) (*entities.Order, error)
}

// InitOrderService initialize the order service with necessary dependencies
func InitOrderService(orderRepository repositories.OrderRepository, stripeAPIKey string) OrderService {
	return &orderService{orderRepository: orderRepository, stripeAPIKey: stripeAPIKey}
}

// GetOrder returns an order by id
func (o orderService) GetOrder(ctx context.Context, id int64) (*entities.Order, error) {
	order, err := o.orderRepository.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	// Retrieve order items
	order.Items, err = o.orderRepository.GetOrderItems(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrders returns all orders from the database
func (o orderService) GetOrders(ctx context.Context) ([]*entities.Order, error) {
	// Retrieve orders
	orders, err := o.orderRepository.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	// Retrieve order items
	for _, order := range orders {
		order.Items, err = o.orderRepository.GetOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (o orderService) GetOrdersByUser(ctx context.Context, userID int64) ([]*entities.Order, error) {
	// Retrieve orders
	orders, err := o.orderRepository.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Retrieve order items
	for _, order := range orders {
		order.Items, err = o.orderRepository.GetOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

// CreateOrder creates a new order and returns it.
func (o orderService) CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	return o.orderRepository.CreateOrder(ctx, order)
}

// UpdateOrder updates an order
// The current implementation is not atomic, it is possible that the order is updated by another process
// while the user is paying for it. Moreover, it is not possible to modify the order items.
func (o orderService) UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	return order, o.orderRepository.UpdateOrder(ctx, order)
}

// PayOrder is a method that performs the payment of an order with the default method (credit card).
// The payment is performed using the Stripe API. In this implementation, the payment is performed using the charge
// API. In Europe, the payment should be performed using the Intents API (see https://stripe.com/docs/payments/intents)
// and the related documentation to migrate to the Intents API (see https://stripe.com/docs/payments/payment-intents/migration/charges).
// It returns the updated order with the payment id and the payment status
func (o orderService) PayOrder(ctx context.Context, id int64, user *entities.User) (*entities.Order, error) {
	// Get the order
	order, err := o.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	// Set Stripe API key
	stripe.Key = o.stripeAPIKey

	// Attempt to make the charge.
	// We are setting the charge response to _
	// as we are not using it.
	c, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(int64(order.TotalPrice * 100)),
		Currency:     stripe.String(string(stripe.CurrencyEUR)),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")}, // this should come from clientside
		ReceiptEmail: stripe.String(user.Email)})
	if err != nil {
		return nil, err
	}

	// Set the payment status to paid
	order.Status = entities.OrderStatusPaid
	order.PaymentID = c.ID

	// Update the order
	order, err = o.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}