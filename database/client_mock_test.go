package database_test

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
	"github.com/josephspurrier/gocleanarchitecture/domain/user"
)

// TestMockService ensures the mock service works correctly.
func TestMockService(t *testing.T) {
	// Test the reading and writing.
	s := new(database.MockService)
	AssertEqual(t, s.Read(), nil)
	AssertEqual(t, s.Write(), nil)

	// Test adding a record and reading it.
	u := new(user.Item)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	s.AddRecord(*u)
	AssertEqual(t, len(s.Records()), 1)
}
