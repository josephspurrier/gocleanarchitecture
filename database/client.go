package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// Schema represents the database structure.
type Schema struct {
	Records []domain.User
}

// Client represents a client to the data store.
type Client struct {
	// Path is the relative filename.
	Path string

	data  *Schema
	mutex sync.RWMutex
}

// Service represents a service for interacting with the database.
type Service interface {
	Read() error
	Write() error
	Records() []domain.User
	AddRecord(domain.User)
}

// NewClient returns a new database client.
func NewClient(path string) *Client {
	c := &Client{
		Path: path,
		data: new(Schema),
	}

	return c
}

// Reads opens/initializes the database.
func (c *Client) Read() error {
	var err error
	var b []byte

	c.mutex.Lock()

	if _, err = os.Stat(c.Path); os.IsNotExist(err) {
		err = ioutil.WriteFile(c.Path, []byte("{}"), 0644)
		if err != nil {
			c.mutex.Unlock()
			return err
		}
	}

	b, err = ioutil.ReadFile(c.Path)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	c.data = new(Schema)
	err = json.Unmarshal(b, &c.data)

	c.mutex.Unlock()

	return err
}

// Write saves the database.
func (c *Client) Write() error {
	var err error
	var b []byte

	c.mutex.Lock()

	b, err = json.Marshal(c.data)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	err = ioutil.WriteFile(c.Path, b, 0644)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	c.mutex.Unlock()

	return err
}

// AddRecord adds a record to the database.
func (c *Client) AddRecord(rec domain.User) {
	c.data.Records = append(c.data.Records, rec)
}

// Records retrieves all records from the database.
func (c *Client) Records() []domain.User {
	return c.data.Records
}
