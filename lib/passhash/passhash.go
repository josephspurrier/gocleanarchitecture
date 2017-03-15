// Package passhash provides password hashing functionality using bcrypt.
package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

// Item represents a password hashing system.
type Item struct{}

// Hash returns a hashed string and an error.
func (s *Item) Hash(password string) (string, error) {
	return HashString(password)
}

// Match returns true if the hash matches the password.
func (s *Item) Match(hash, password string) bool {
	return MatchString(hash, password)
}

// HashString returns a hashed string and an error.
func HashString(password string) (string, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(key), err
}

// HashBytes returns a hashed byte array and an error.
func HashBytes(password []byte) ([]byte, error) {
	key, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return key, err
}

// MatchString returns true if the hash matches the password.
func MatchString(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}

	return false
}

// MatchBytes returns true if the hash matches the password.
func MatchBytes(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == nil {
		return true
	}

	return false
}
