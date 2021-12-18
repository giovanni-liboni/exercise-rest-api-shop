package entities

type Order struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Payment_method string `json:"payment_method"`
	PaymentID int64 `json:"payment_id"`
	Total_price float64 `json:"total_price"`
	Status string `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
