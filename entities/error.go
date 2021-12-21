package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInvalidUsername   = errors.New("Invalid username")
	ErrEmptyPassword     = errors.New("Empty password")
)
