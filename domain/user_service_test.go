package domain_test

import (
	"errors"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/adapter"
	"github.com/josephspurrier/gocleanarchitecture/adapter/jsonrepo"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"

	"github.com/stretchr/testify/assert"
)

//  BadHasher represents a password hashing system that always fails.
type BadHasher struct{}

// Hash returns a hashed string and an error.
func (s *BadHasher) Hash(password string) (string, error) {
	return "", errors.New("Forced error.")
}

// Match returns true if the hash matches the password.
func (s *BadHasher) Match(hash, password string) bool {
	return false
}

// setup handles the creation of the service.
func setup() *domain.UserService {
	return domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(adapter.Passhash))
}

// TestCreateUser ensures user can be created.
func TestCreateUser(t *testing.T) {
	// Test user creation.
	s := setup()

	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Create(u)
	assert.Equal(t, err, nil)

	// Test user creation fail.
	err = s.Create(u)
	assert.Equal(t, err, domain.ErrUserAlreadyExist)

	// Test user query.
	uTest, err := s.ByEmail("jdoe@example.com")
	assert.Equal(t, err, nil)
	assert.Equal(t, uTest.Email, "jdoe@example.com")

	// Test failed user query.
	_, err = s.ByEmail("bademail@example.com")
	assert.Equal(t, err, domain.ErrUserNotFound)
}

// TestAuthenticate ensures user can authenticate.
func TestAuthenticate(t *testing.T) {
	// Test user creation.
	s := setup()

	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Create(u)
	assert.Equal(t, err, nil)

	// Test user authentication.
	err = s.Authenticate(u)
	assert.Equal(t, err, nil)

	// Test failed user authentication.
	u.Password = "BadPa$$w0rd"
	err = s.Authenticate(u)
	assert.Equal(t, err, domain.ErrUserPasswordNotMatch)

	// Test failed user authentication.
	u.Email = "bfranklin@example.com"
	err = s.Authenticate(u)
	assert.Equal(t, err, domain.ErrUserNotFound)
}

// TestUserFailures ensures user fails properly.
func TestUserFailures(t *testing.T) {
	// Test user creation.
	db := new(jsondb.MockService)
	s := domain.NewUserService(jsonrepo.NewUserRepo(db), new(adapter.Passhash))

	db.WriteFail = true
	db.ReadFail = true

	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Create(u)
	assert.NotNil(t, err)

	// Test user authentication.
	err = s.Authenticate(u)
	assert.NotNil(t, err)

	// Test failed user query.
	_, err = s.ByEmail("favalon@example.com")
	assert.NotNil(t, err)

	// Test failed user authentication.
	u.Email = "bfranklin@example.com"
	err = s.Authenticate(u)
	assert.NotNil(t, err)
}

// TestBadHasherFailures ensures user fails properly.
func TestBadHasherFailures(t *testing.T) {
	// Test user creation.
	db := new(jsondb.MockService)
	s := domain.NewUserService(jsonrepo.NewUserRepo(db), new(BadHasher))

	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Create(u)
	assert.Equal(t, err, domain.ErrPasswordHash)
}
