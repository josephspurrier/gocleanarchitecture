package repository_test

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/repository"
)

// TestUserRepo tests the user repo.
func TestUserRepo(t *testing.T) {
	db := new(repository.MockService)
	s := repository.NewUserRepo(db)

	_, err := s.FindByEmail("bad@example.com")
	AssertEqual(t, err, domain.ErrUserNotFound)

	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err = s.Store(u)
	AssertEqual(t, err, nil)

	_, err = s.FindByEmail("jdoe@example.com")
	AssertEqual(t, err, nil)
}

// TestUserRepoFail tests the user repo.
func TestUserRepoFail(t *testing.T) {
	db := new(repository.MockService)
	s := repository.NewUserRepo(db)

	db.WriteFail = true
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Store(u)
	AssertNotNil(t, err)

	db.WriteFail = false
	err = s.Store(u)
	AssertEqual(t, err, nil)

	_, err = s.FindByEmail("jdoe@example.com")
	AssertEqual(t, err, nil)

	db.ReadFail = true
	_, err = s.FindByEmail("jdoe@example.com")
	AssertNotNil(t, err)

	err = s.Store(u)
	AssertNotNil(t, err)
}
