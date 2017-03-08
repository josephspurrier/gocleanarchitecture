package database_test

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// TestCreateUser ensures user can be created.
func TestCreateUser(t *testing.T) {
	// Test user creation.
	db := new(database.MockService)
	s := database.NewUserService(db)
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
	AssertEqual(t, uTest.Password, "Pa$$w0rd")

	// Test failed user query.
	_, err = s.User("bademail@example.com")
	AssertEqual(t, err, domain.ErrUserNotFound)
}

// TestAuthenticate ensures user can authenticate.
func TestAuthenticate(t *testing.T) {
	// Test user creation.
	db := new(database.MockService)
	s := database.NewUserService(db)
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

// TestUserFailures ensures fails properly.
func TestUserFailures(t *testing.T) {
	// Test user creation.
	db := new(database.MockService)
	s := database.NewUserService(db)

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
