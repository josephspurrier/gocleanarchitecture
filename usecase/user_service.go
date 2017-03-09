package usecase

import "github.com/josephspurrier/gocleanarchitecture/domain"

// UserService represents a service for managing users.
type UserService struct {
	userRepo domain.UserRepo
}

// NewUserServer returns the service for managing users.
func NewUserService(repo domain.UserRepo) *UserService {
	s := new(UserService)
	s.userRepo = repo
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(item *domain.User) error {
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
func (s *UserService) User(email string) (*domain.User, error) {
	item, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return item, domain.ErrUserNotFound
	}

	return item, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(item *domain.User) error {
	_, err := s.userRepo.FindByEmail(item.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	return s.userRepo.Store(item)
}
