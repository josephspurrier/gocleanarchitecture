package jsondb

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"

	"github.com/stretchr/testify/assert"
)

// TestClient ensures the client works properly.
func TestClient(t *testing.T) {
	c := New("testdata")

	recordType := "bar"

	// Check the output.
	assert.Equal(t, c.Path, "testdata")
	assert.Equal(t, c.write(recordType), nil)
	assert.Equal(t, c.read(recordType), nil)
	assert.Equal(t, c.write(recordType), nil)

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	c.AddRecord(recordType, *u)
	records, err := c.Records(recordType)
	assert.Equal(t, len(records), 1)
	assert.Equal(t, err, nil)

	// Cleanup
	err = os.Remove(path.Join("testdata", recordType+".json"))
	if err != nil {
		t.Error(err)
	}
}

// TestClientFailOpen ensures the client fails properly.
func TestClientFailOpen(t *testing.T) {
	c := New("testdata")

	filePath := path.Join("testdata", "foobad.json")

	// Write a bad file.
	err := ioutil.WriteFile(filePath, []byte("{"), 0644)
	if err != nil {
		t.Error(err)
	}

	// Check the output.
	assert.NotNil(t, c.read("foobad"))

	// Cleanup
	err = os.Remove(filePath)
	if err != nil {
		t.Error(err)
	}
}
