package database_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/database"
)

// AssertEqual throws an error if the two values are not equal.
func AssertEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}

// AssertNotNil throws an error if the value is nil.
func AssertNotNil(t *testing.T, actualValue interface{}) {
	if actualValue == nil {
		t.Errorf("\n got: %v\ndidn't want: %v", actualValue, nil)
	}
}

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
