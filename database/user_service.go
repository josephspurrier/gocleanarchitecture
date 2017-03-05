package database

import (
	"errors"
	"strings"

	"github.com/josephspurrier/gocleanarchitecture/domain/user"
)

var (
	// ErrNotFound is when the user does not exist.
	ErrNotFound = errors.New("User not found.")
	// ErrAlreadyExist is when the user already exists.
	ErrAlreadyExist = errors.New("User already exists.")
	// ErrPasswordNotMatch is when the user's password does not match.
	ErrPasswordNotMatch = errors.New("User password does not match.")
)

// UserService represents a service for managing users.
type UserService struct {
	client *Client
}

// Autheticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(d *user.Item) error {
	// Load the data.
	err := s.client.Read()
	if err != nil {
		return err
	}

	// Determine if the record exists.
	for _, v := range s.client.data.Records {
		if v.Email == d.Email {
			if v.Password == d.Password {
				return nil
			}
			return ErrPasswordNotMatch
		}
	}

	return ErrNotFound
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
	for _, v := range s.client.data.Records {
		if v.Email == email {
			// Return the record.
			return &v, nil
		}
	}

	return item, ErrNotFound
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(d *user.Item) error {
	err := s.client.Read()
	if err != nil {
		return err
	}

	// Check if the user already exists
	for _, v := range s.client.data.Records {
		if strings.ToLower(v.Email) == strings.ToLower(d.Email) {
			// Return an error since the record exists.
			return ErrAlreadyExist
		}
	}

	// Add the record.
	s.client.data.Records = append(s.client.data.Records, *d)

	// Save the record to the database.
	return s.client.Write()
}
