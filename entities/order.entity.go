package entities

import "time"

type Order struct {
	ID             int       `json:"id" db:"id"`
	UserID         int       `json:"user_id" db:"user_id"`
	Payment_method string    `json:"payment_method" db:"payment_method"`
	PaymentID      int       `json:"payment_id" db:"payment_id"`
	Total_price    float32   `json:"total_price" db:"total_price"`
	Status         string    `json:"status" db:"status"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
