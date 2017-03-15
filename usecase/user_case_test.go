package usecase_test

import (
	"errors"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/passhash"
	"github.com/josephspurrier/gocleanarchitecture/repository"
	"github.com/josephspurrier/gocleanarchitecture/usecase"
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

// TestCreateUser ensures user can be created.
func TestCreateUser(t *testing.T) {
	// Test user creation.
	s := usecase.NewUserCase(repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err := s.CreateUser(u)
	AssertEqual(t, err, nil)

	// Test user creation fail.
	err = s.CreateUser(u)
	AssertEqual(t, err, domain.ErrUserAlreadyExist)

	// Test user query.
	uTest, err := s.User("jdoe@example.com")
	AssertEqual(t, err, nil)
	AssertEqual(t, uTest.Email, "jdoe@example.com")

	// Test failed user query.
	_, err = s.User("bademail@example.com")
	AssertEqual(t, err, domain.ErrUserNotFound)
}

// TestAuthenticate ensures user can authenticate.
func TestAuthenticate(t *testing.T) {
	// Test user creation.
	s := usecase.NewUserCase(repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.CreateUser(u)
	AssertEqual(t, err, nil)

	// Test user authentication.
	err = s.Authenticate(u)
	AssertEqual(t, err, nil)

	// Test failed user authentication.
	u.Password = "BadPa$$w0rd"
	err = s.Authenticate(u)
	AssertEqual(t, err, domain.ErrUserPasswordNotMatch)

	// Test failed user authentication.
	u.Email = "bfranklin@example.com"
	err = s.Authenticate(u)
	AssertEqual(t, err, domain.ErrUserNotFound)
}

// TestUserFailures ensures user fails properly.
func TestUserFailures(t *testing.T) {
	// Test user creation.
	db := new(repository.MockService)
	s := usecase.NewUserCase(repository.NewUserRepo(db), new(passhash.Item))

	db.WriteFail = true
	db.ReadFail = true

	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.CreateUser(u)
	AssertNotNil(t, err)

	// Test user authentication.
	err = s.Authenticate(u)
	AssertNotNil(t, err)

	// Test failed user query.
	_, err = s.User("favalon@example.com")
	AssertNotNil(t, err)

	// Test failed user authentication.
	u.Email = "bfranklin@example.com"
	err = s.Authenticate(u)
	AssertNotNil(t, err)
}

// TestBadHasherFailures ensures user fails properly.
func TestBadHasherFailures(t *testing.T) {
	// Test user creation.
	db := new(repository.MockService)
	s := usecase.NewUserCase(repository.NewUserRepo(db), new(BadHasher))

	u := new(domain.User)
	u.Email = "ssmith@example.com"
	u.Password = "Pa$$w0rd"
	err := s.CreateUser(u)
	AssertEqual(t, err, domain.ErrPasswordHash)
}
