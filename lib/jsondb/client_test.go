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
	AssertEqual(t, c.write(), nil)
	AssertEqual(t, c.read(), nil)
	AssertEqual(t, c.write(), nil)

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	c.AddRecord(*u)
	records, err := c.Records()
	AssertEqual(t, len(records), 1)
	AssertEqual(t, err, nil)

	// Cleanup
	err = os.Remove("db.json")
	if err != nil {
		t.Error(err)
	}
}

// TestClient ensures the client fails properly.
func TestClientFail(t *testing.T) {
	c := NewClient("")

	// Check the output.
	AssertEqual(t, c.Path, "")
	AssertNotNil(t, c.write())
	AssertNotNil(t, c.read())
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
	AssertNotNil(t, c.read())

	// Cleanup
	err = os.Remove("dbbad.json")
	if err != nil {
		t.Error(err)
	}
}
