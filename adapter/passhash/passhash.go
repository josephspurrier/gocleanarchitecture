package passhash

import (
	phl "github.com/josephspurrier/gocleanarchitecture/lib/passhash"
)

// Item implements the password hashing system.
type Item struct{}

// Hash returns a hashed string or an error.
func (s *Item) Hash(password string) (string, error) {
	return phl.HashString(password)
}

// Match returns true if the hash matches the password.
func (s *Item) Match(hash, password string) bool {
	return phl.MatchString(hash, password)
}
