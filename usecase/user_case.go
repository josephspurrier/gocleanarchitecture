package usecase

import "github.com/josephspurrier/gocleanarchitecture/domain"

// UserCase represents a service for managing users.
type UserCase struct {
	userRepo domain.UserRepo
	passhash domain.PasshashCase
}

// NewUserCase returns the service for managing users.
func NewUserCase(repo domain.UserRepo, passhash domain.PasshashCase) *UserCase {
	s := new(UserCase)
	s.userRepo = repo
	s.passhash = passhash
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserCase) Authenticate(item *domain.User) error {
	q, err := s.userRepo.FindByEmail(item.Email)
	if err != nil {
		return domain.ErrUserNotFound
	}

	// If passwords match.
	if s.passhash.Match(q.Password, item.Password) {
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

	passNew, err := s.passhash.Hash(item.Password)
	if err != nil {
		return domain.ErrPasswordHash
	}

	// Swap the password.
	passOld := item.Password
	item.Password = passNew
	err = s.userRepo.Store(item)

	// Restore the password.
	item.Password = passOld

	return err
}
