package domain

import (
	"errors"
)

var (
	// ErrPasswordHash
	ErrPasswordHash = errors.New("Password hash failed.")
)

// Passhash represents a password hashing system.
type Passhash struct {
}

// PasshashCase represents a service for managing hashed passwords.
type PasshashCase interface {
	Hash(password string) (string, error)
	Match(hash, password string) bool
}
