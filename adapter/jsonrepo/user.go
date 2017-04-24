package jsonrepo

import (
	"github.com/josephspurrier/gocleanarchitecture/domain"

	"github.com/mitchellh/mapstructure"
)

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

	// Retrieve the records
	records, err := s.client.Records("user")
	if err != nil {
		return item, err
	}

	// Determine if the record exists.
	for _, r := range records {
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
	// Add the record.
	return s.client.AddRecord("user", *item)
}
