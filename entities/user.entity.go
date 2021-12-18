package entities

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-" db:"password"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Role string `json:"role"`
}