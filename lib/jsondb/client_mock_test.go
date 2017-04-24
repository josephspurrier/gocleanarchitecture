package jsondb

import (
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"

	"github.com/stretchr/testify/assert"
)

// TestMockService ensures the mock service works correctly.
func TestMockService(t *testing.T) {
	// Test the reading and writing.
	s := new(MockService)

	// Test forced failures.
	s.WriteFail = true
	s.ReadFail = true
	assert.NotNil(t, s.read())
	assert.NotNil(t, s.write())

	// Test no failures.
	s.WriteFail = false
	s.ReadFail = false
	assert.Equal(t, nil, s.read())
	assert.Equal(t, nil, s.write())

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	s.AddRecord("user", *u)
	records, err := s.Records("user")
	assert.Equal(t, 1, len(records))
	assert.Equal(t, nil, err)
}
