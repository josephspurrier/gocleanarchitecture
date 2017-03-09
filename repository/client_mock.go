package repository

import (
	"errors"
	"sync"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// MockService represents a service for managing users.
type MockService struct {
	records   []domain.User
	mutex     sync.RWMutex
	ReadFail  bool
	WriteFail bool
}

// Reads reads database.
func (c *MockService) Read() error {
	if c.ReadFail {
		return errors.New("Read failure.")
	}
	return nil
}

// Write saves the database.
func (c *MockService) Write() error {
	if c.WriteFail {
		return errors.New("Write failure.")
	}
	return nil
}

// AddRecord adds a record to the database.
func (c *MockService) AddRecord(rec domain.User) {
	c.mutex.Lock()
	c.records = append(c.records, rec)
	c.mutex.Unlock()
}

// Records retrieves all records from the database.
func (c *MockService) Records() []domain.User {
	c.mutex.RLock()
	r := c.records
	c.mutex.RUnlock()
	return r
}
