package jsonrepo_test

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/adapter/jsonrepo"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"

	"github.com/stretchr/testify/assert"
)

// TestUserRepo tests the user repo.
func TestUserRepo(t *testing.T) {
	db := new(jsondb.MockService)
	s := jsonrepo.NewUserRepo(db)

	_, err := s.ByEmail("bad@example.com")
	assert.Equal(t, err, domain.ErrUserNotFound)

	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err = s.Store(u)
	assert.Equal(t, err, nil)

	_, err = s.ByEmail("jdoe@example.com")
	assert.Equal(t, err, nil)
}

// TestUserRepoFail tests the user repo.
func TestUserRepoFail(t *testing.T) {
	db := new(jsondb.MockService)
	s := jsonrepo.NewUserRepo(db)

	db.WriteFail = true
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	err := s.Store(u)
	assert.NotNil(t, err)

	db.WriteFail = false
	err = s.Store(u)
	assert.Equal(t, err, nil)

	_, err = s.ByEmail("jdoe@example.com")
	assert.Equal(t, err, nil)

	db.ReadFail = true
	_, err = s.ByEmail("jdoe@example.com")
	assert.NotNil(t, err)

	err = s.Store(u)
	assert.NotNil(t, err)
}
