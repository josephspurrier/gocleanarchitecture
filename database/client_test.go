package database_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
	"github.com/josephspurrier/gocleanarchitecture/domain/user"
)

// TestClient ensures the client works properly.
func TestClient(t *testing.T) {
	c := database.NewClient("db.json")

	// Check the output.
	AssertEqual(t, c.Path, "db.json")
	AssertEqual(t, c.Write(), nil)
	AssertEqual(t, c.Read(), nil)
	AssertEqual(t, c.Write(), nil)

	// Test adding a record and reading it.
	u := new(user.Item)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	c.AddRecord(*u)
	AssertEqual(t, len(c.Records()), 1)

	// Cleanup
	os.Remove("db.json")
}
