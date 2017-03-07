package database_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
)

// TestClient ensures the client works properly.
func TestClient(t *testing.T) {
	c := database.NewClient("db.json")

	// Check the output.
	AssertEqual(t, c.Path, "db.json")
	AssertEqual(t, c.Write(), nil)
	AssertEqual(t, c.Read(), nil)
	AssertEqual(t, c.Write(), nil)

	// Cleanup
	os.Remove("db.json")
}
