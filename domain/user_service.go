package domain

// IUserService is the interface for users.
type IUserService interface {
	ByEmail(email string) (*User, error)
	Create(firstname, lastname, email, password string) error
	Authenticate(email, password string) error
}

// UserService implements the service for users.
type UserService struct {
	repo IUserRepo // Storage.
	hash IPasshash // Password hashing.
}

// IUserRepo represents the service for storage of users.
type IUserRepo interface {
	ByEmail(email string) (*User, error)
	Store(item *User) error
}

// NewUserService returns the service for managing users.
func NewUserService(repo IUserRepo, hash IPasshash) *UserService {
	s := new(UserService)
	s.repo = repo
	s.hash = hash
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserService) Authenticate(email, password string) error {
	q, err := s.repo.ByEmail(email)
	if err != nil {
		return ErrUserNotFound
	}

	// If passwords match.
	if s.hash.Match(q.Password, password) {
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
func (s *UserService) Create(firstname, lastname, email,
	password string) error {
	// Check if a user already exists.
	_, err := s.repo.ByEmail(email)
	if err == nil {
		return ErrUserAlreadyExist
	}

	// Hash the password.
	passNew, err := s.hash.Hash(password)
	if err != nil {
		return ErrPasswordHash
	}

	// Create the user.
	item := new(User)
	item.FirstName = firstname
	item.LastName = lastname
	item.Email = email
	item.Password = passNew

	// Store the user.
	err = s.repo.Store(item)

	return err
}
