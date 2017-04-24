package domain

import "errors"

var (
	// ErrPasswordHash is when a password hash creation operation fails.
	ErrPasswordHash = errors.New("password hash failed")
)

// IPasshash is the interface for password hashing.
type IPasshash interface {
	Hash(password string) (string, error)
	Match(hash, password string) bool
}
