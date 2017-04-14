package domain

// IUserService is the interface for users.
type IUserService interface {
	ByEmail(email string) (*User, error)
	Create(item *User) error
	Authenticate(item *User) error
}

// UserService implements the service for users.
type UserService struct {
	repo UserRepo  // Storage.
	hash IPasshash // Password hashing.
}

// UserRepo represents the service for storage of users.
type UserRepo interface {
	ByEmail(email string) (*User, error)
	Store(item *User) error
}

// NewUserService returns the service for managing users.
func NewUserService(repo UserRepo, hash IPasshash) *UserService {
	s := new(UserService)
	s.repo = repo
	s.hash = hash
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(item *User) error {
	q, err := s.repo.ByEmail(item.Email)
	if err != nil {
		return ErrUserNotFound
	}

	// If passwords match.
	if s.hash.Match(q.Password, item.Password) {
		return nil
	}

	return ErrUserPasswordNotMatch
}

// ByEmail returns a user by email or an error if not found.
func (s *UserService) ByEmail(email string) (*User, error) {
	item, err := s.repo.ByEmail(email)
	if err != nil {
		return item, ErrUserNotFound
	}

	return item, nil
}

// Create a new user.
func (s *UserService) Create(item *User) error {
	_, err := s.repo.ByEmail(item.Email)
	if err == nil {
		return ErrUserAlreadyExist
	}

	passNew, err := s.hash.Hash(item.Password)
	if err != nil {
		return ErrPasswordHash
	}

	// Swap the password.
	passOld := item.Password
	item.Password = passNew
	err = s.repo.Store(item)

	// Restore the password.
	item.Password = passOld

	return err
}
