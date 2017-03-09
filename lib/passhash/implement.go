package passhash

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
