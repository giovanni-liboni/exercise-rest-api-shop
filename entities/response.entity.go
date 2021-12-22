package entities

type Response struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type AppError struct {
	Err        error
	StatusCode int
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}

type AppResult struct {
	Data       interface{}
	Message    string
	Err        error
	StatusCode int
}
