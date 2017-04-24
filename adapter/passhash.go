package adapter

import (
	"github.com/josephspurrier/gocleanarchitecture/lib/passhash"
)

// Passhash implements the password hashing system.
type Passhash struct{}

// Hash returns a hashed string or an error.
func (s *Passhash) Hash(password string) (string, error) {
	return passhash.HashString(password)
}

// Match returns true if the hash matches the password.
func (s *Passhash) Match(hash, password string) bool {
	return passhash.MatchString(hash, password)
}
