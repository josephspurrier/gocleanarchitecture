package domain

import "errors"

var (
	// ErrUserNotFound is when the user does not exist.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExist is when the user already exists.
	ErrUserAlreadyExist = errors.New("user already exists")
	// ErrUserPasswordNotMatch is when the user's password does not match.
	ErrUserPasswordNotMatch = errors.New("user password does not match")
)

// User represents a user of the system.
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}
