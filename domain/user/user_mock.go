package user

import (
	"strings"
	"sync"
)

// MockService represents a service for managing users.
type MockService struct {
	records []Item
	mutex   sync.RWMutex
}

// Authenticate returns an error if the email and password don't match.
func (s *MockService) Authenticate(d *Item) error {
	s.mutex.RLock()

	// Determine if the record exists.
	for _, v := range s.records {
		if v.Email == d.Email {
			if v.Password == d.Password {
				s.mutex.RUnlock()
				return nil
			}
			s.mutex.RUnlock()
			return ErrPasswordNotMatch
		}
	}

	s.mutex.RUnlock()

	return ErrNotFound
}

// User returns a user by email.
func (s *MockService) User(email string) (*Item, error) {
	item := new(Item)

	s.mutex.RLock()

	// Determine if the record exists.
	for _, v := range s.records {
		if v.Email == email {
			// Return the record.
			s.mutex.RUnlock()
			return &v, nil
		}
	}

	s.mutex.RUnlock()

	return item, ErrNotFound
}

// CreateUser creates a new user.
func (s *MockService) CreateUser(d *Item) error {
	s.mutex.Lock()

	// Check if the user already exists
	for _, v := range s.records {
		if strings.ToLower(v.Email) == strings.ToLower(d.Email) {
			// Return an error since the record exists.
			s.mutex.Unlock()
			return ErrAlreadyExist
		}
	}

	// Add the record.
	s.records = append(s.records, *d)

	s.mutex.Unlock()

	// Save the record to the database.
	return nil
}
