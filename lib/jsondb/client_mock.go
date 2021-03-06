package jsondb

import (
	"errors"
	"sync"
)

var (
	// ErrMockRead is a read failure.
	ErrMockRead = errors.New("read failure")
	// ErrMockWrite is a write failure.
	ErrMockWrite = errors.New("write failure")
)

// MockService represents a service for managing users.
type MockService struct {
	records   []interface{}
	mutex     sync.RWMutex
	ReadFail  bool
	WriteFail bool
}

// Reads reads database.
func (c *MockService) read() error {
	if c.ReadFail {
		return ErrMockRead
	}
	return nil
}

// Write saves the database.
func (c *MockService) write() error {
	if c.WriteFail {
		return ErrMockWrite
	}
	return nil
}

// AddRecord adds a record to the database.
func (c *MockService) AddRecord(recordType string, rec interface{}) error {
	err := c.read()
	if err != nil {
		return err
	}

	c.mutex.Lock()
	c.records = append(c.records, rec)
	c.mutex.Unlock()

	return c.write()
}

// Records retrieves all records from the database.
func (c *MockService) Records(recordType string) ([]interface{}, error) {
	err := c.read()
	if err != nil {
		return c.records, err
	}

	c.mutex.RLock()
	r := c.records
	c.mutex.RUnlock()

	return r, nil
}
