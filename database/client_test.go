package database_test

import (
	"io/ioutil"
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

// TestClient ensures the client fails properly.
func TestClientFail(t *testing.T) {
	c := database.NewClient("")

	// Check the output.
	AssertEqual(t, c.Path, "")
	AssertNotNil(t, c.Write())
	AssertNotNil(t, c.Read())
}

// TestClientFailOpen ensures the client fails properly.
func TestClientFailOpen(t *testing.T) {
	c := database.NewClient("dbbad.json")

	// Write a bad file.
	ioutil.WriteFile("dbbad.json", []byte("{"), 0644)

	// Check the output.
	AssertNotNil(t, c.Read())

	// Cleanup
	os.Remove("dbbad.json")
}
