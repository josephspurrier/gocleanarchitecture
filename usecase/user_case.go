package usecase

import "github.com/josephspurrier/gocleanarchitecture/domain"

// UserCase represents a service for managing users.
type UserCase struct {
	userRepo domain.UserRepo
}

// NewUserCase returns the service for managing users.
func NewUserCase(repo domain.UserRepo) *UserCase {
	s := new(UserCase)
	s.userRepo = repo
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserCase) Authenticate(item *domain.User) error {
	q, err := s.userRepo.FindByEmail(item.Email)
	if err != nil {
		return domain.ErrUserNotFound
	}

	// If passwords match.
	if q.Password == item.Password {
		return nil
	}

	return domain.ErrUserPasswordNotMatch
}

// User returns a user by email.
func (s *UserCase) User(email string) (*domain.User, error) {
	item, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return item, domain.ErrUserNotFound
	}

	return item, nil
}

// CreateUser creates a new user.
func (s *UserCase) CreateUser(item *domain.User) error {
	_, err := s.userRepo.FindByEmail(item.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	return s.userRepo.Store(item)
}
