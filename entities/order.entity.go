package entities

import "time"

const (
	OrderStatusCreated  = "created"
	OrderStatusCanceled = "canceled"
	OrderStatusPending  = "pending"
	OrderStatusPaid     = "paid"
)

type Order struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentID     string    `json:"payment_id" db:"payment_id"`
	TotalPrice    float32   `json:"total_price" db:"total_price"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Items         []*Item   `json:"items"`
}
