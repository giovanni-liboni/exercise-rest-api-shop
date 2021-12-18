package entities

import "time"

type Order struct {
	ID             int     `json:"id"`
	UserID         int     `json:"user_id"`
	Payment_method string    `json:"payment_method"`
	PaymentID      int     `json:"payment_id"`
	Total_price    float32   `json:"total_price"`
	Status         string    `json:"status"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
