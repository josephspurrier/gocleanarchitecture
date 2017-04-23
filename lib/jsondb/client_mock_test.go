package jsondb

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// TestMockService ensures the mock service works correctly.
func TestMockService(t *testing.T) {
	// Test the reading and writing.
	s := new(MockService)

	// Test forced failures.
	s.WriteFail = true
	s.ReadFail = true
	AssertNotNil(t, s.read())
	AssertNotNil(t, s.write())

	// Test no failures.
	s.WriteFail = false
	s.ReadFail = false
	AssertEqual(t, s.read(), nil)
	AssertEqual(t, s.write(), nil)

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	s.AddRecord(*u)
	records, err := s.Records()
	AssertEqual(t, len(records), 1)
	AssertEqual(t, err, nil)
}
