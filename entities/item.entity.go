package entities

type Item struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Producer   string `json:"producer"`
	Category    string `json:"category"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}