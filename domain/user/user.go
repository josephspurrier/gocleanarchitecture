package user

import "errors"

var (
	// ErrNotFound is when the user does not exist.
	ErrNotFound = errors.New("User not found.")
	// ErrAlreadyExist is when the user already exists.
	ErrAlreadyExist = errors.New("User already exists.")
	// ErrPasswordNotMatch is when the user's password does not match.
	ErrPasswordNotMatch = errors.New("User password does not match.")
)

// Item represents a user.
type Item struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Service represents a service for managing users.
type Service interface {
	User(email string) (*Item, error)
	CreateUser(user *Item) error
	Authenticate(user *Item) error
}
