package repo

import (
	"github.com/josephspurrier/gocleanarchitecture/domain"

	"github.com/mitchellh/mapstructure"
)

// IRepoService is the interface for storage.
type IRepoService interface {
	Read() error
	Write() error
	Records() []interface{}
	AddRecord(interface{})
}

// UserService implements the service for storage of users.
type UserService struct {
	client IRepoService
}

// NewUserRepo returns the service for storage of users.
func NewUserRepo(client IRepoService) *UserService {
	s := new(UserService)
	s.client = client
	return s
}

// ByEmail returns a user by an email or an error if not found.
func (s *UserService) ByEmail(email string) (*domain.User, error) {
	item := new(domain.User)

	// Load the data.
	err := s.client.Read()
	if err != nil {
		return item, err
	}

	// Determine if the record exists.
	for _, r := range s.client.Records() {
		// Decode the user from the database.
		var v domain.User
		err = mapstructure.Decode(r, &v)
		if err != nil {
			return &v, err
		}

		// Ensure the user is returned.
		if v.Email == email {
			return &v, nil
		}
	}

	return item, domain.ErrUserNotFound
}

// Store adds a user or returns an error.
func (s *UserService) Store(item *domain.User) error {
	// Load the data.
	err := s.client.Read()
	if err != nil {
		return err
	}

	// Add the record.
	s.client.AddRecord(*item)

	// Save the record to the database.
	return s.client.Write()
}
