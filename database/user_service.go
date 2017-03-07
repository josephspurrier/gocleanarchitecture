package database

import (
	"strings"

	"github.com/josephspurrier/gocleanarchitecture/domain/user"
)

// UserService represents a service for managing users.
type UserService struct {
	client Service
}

func NewUserService(client Service) *UserService {
	s := new(UserService)
	s.client = client
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(d *user.Item) error {
	// Load the data.
	err := s.client.Read()
	if err != nil {
		return err
	}

	// Determine if the record exists.
	for _, v := range s.client.Records() {
		if v.Email == d.Email {
			if v.Password == d.Password {
				return nil
			}
			return user.ErrPasswordNotMatch
		}
	}

	return user.ErrNotFound
}

// User returns a user by email.
func (s *UserService) User(email string) (*user.Item, error) {
	item := new(user.Item)

	// Load the data.
	err := s.client.Read()
	if err != nil {
		return item, err
	}

	// Determine if the record exists.
	for _, v := range s.client.Records() {
		if v.Email == email {
			// Return the record.
			return &v, nil
		}
	}

	return item, user.ErrNotFound
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(d *user.Item) error {
	err := s.client.Read()
	if err != nil {
		return err
	}

	// Check if the user already exists
	for _, v := range s.client.Records() {
		if strings.ToLower(v.Email) == strings.ToLower(d.Email) {
			// Return an error since the record exists.
			return user.ErrAlreadyExist
		}
	}

	// Add the record.
	s.client.AddRecord(*d)

	// Save the record to the database.
	return s.client.Write()
}
