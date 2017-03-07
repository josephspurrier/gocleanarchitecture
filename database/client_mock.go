package database

import (
	"sync"

	"github.com/josephspurrier/gocleanarchitecture/domain/user"
)

// MockService represents a service for managing users.
type MockService struct {
	records []user.Item
	mutex   sync.RWMutex
}

// Reads reads database.
func (c *MockService) Read() error {
	return nil
}

// Write saves the database.
func (c *MockService) Write() error {
	return nil
}

// AddRecord adds a record to the database.
func (c *MockService) AddRecord(rec user.Item) {
	c.mutex.Lock()
	c.records = append(c.records, rec)
	c.mutex.Unlock()
}

// Records retrieves all records from the database.
func (c *MockService) Records() []user.Item {
	c.mutex.RLock()
	r := c.records
	c.mutex.RUnlock()
	return r
}
