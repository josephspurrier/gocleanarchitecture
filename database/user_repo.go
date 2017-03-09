package database

import "github.com/josephspurrier/gocleanarchitecture/domain"

// UserRepo represents a service for storage of users.
type UserRepo struct {
	client Service
}

// NewUserRepo returns the service for storage of users.
func NewUserRepo(client Service) *UserRepo {
	s := new(UserRepo)
	s.client = client
	return s
}

// FindByEmail returns a user by an email.
func (s *UserRepo) FindByEmail(email string) (*domain.User, error) {
	item := new(domain.User)

	// Load the data.
	err := s.client.Read()
	if err != nil {
		return item, err
	}

	// Determine if the record exists.
	for _, v := range s.client.Records() {
		if v.Email == email {
			return &v, nil
		}
	}

	return item, domain.ErrUserNotFound
}

// Store adds a user.
func (s *UserRepo) Store(item *domain.User) error {
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
