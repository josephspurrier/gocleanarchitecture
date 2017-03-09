package usecase

import "github.com/josephspurrier/gocleanarchitecture/domain"

type UserRepo interface {
	Store(item *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type UserService struct {
	UserRepo UserRepo
}

// Authenticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(item *domain.User) error {
	q, err := s.UserRepo.FindByEmail(item.Email)
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
func (s *UserService) User(email string) (*domain.User, error) {
	item, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return item, domain.ErrUserNotFound
	}

	return item, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(item *domain.User) error {
	_, err := s.UserRepo.FindByEmail(item.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	return s.UserRepo.Store(item)
}
