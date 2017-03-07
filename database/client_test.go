package database_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
)

// TestNew ensures the NewClient function works properly.
func TestNewClient(t *testing.T) {
	c := database.NewClient("db.json")

	// Check the output.
	AssertEqual(t, c.Path, "db.json")
	AssertNotNil(t, c.UserService())
}

// TestRead ensures the read and write works properly.
func TestReadWrite(t *testing.T) {
	c := database.NewClient("db.json")

	// Check the output.
	AssertEqual(t, c.Write(), nil)
	AssertEqual(t, c.Read(), nil)
	AssertEqual(t, c.Write(), nil)

	// Cleanup
	os.Remove("db.json")
}
