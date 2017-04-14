package jsondb

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// TestClient ensures the client works properly.
func TestClient(t *testing.T) {
	c := NewClient("db.json")

	// Check the output.
	AssertEqual(t, c.Path, "db.json")
	AssertEqual(t, c.Write(), nil)
	AssertEqual(t, c.Read(), nil)
	AssertEqual(t, c.Write(), nil)

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	c.AddRecord(*u)
	AssertEqual(t, len(c.Records()), 1)

	// Cleanup
	err := os.Remove("db.json")
	if err != nil {
		t.Error(err)
	}
}

// TestClient ensures the client fails properly.
func TestClientFail(t *testing.T) {
	c := NewClient("")

	// Check the output.
	AssertEqual(t, c.Path, "")
	AssertNotNil(t, c.Write())
	AssertNotNil(t, c.Read())
}

// TestClientFailOpen ensures the client fails properly.
func TestClientFailOpen(t *testing.T) {
	c := NewClient("dbbad.json")

	// Write a bad file.
	err := ioutil.WriteFile("dbbad.json", []byte("{"), 0644)
	if err != nil {
		t.Error(err)
	}

	// Check the output.
	AssertNotNil(t, c.Read())

	// Cleanup
	err = os.Remove("dbbad.json")
	if err != nil {
		t.Error(err)
	}
}
