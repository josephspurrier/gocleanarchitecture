package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"gocleanarchitecture/domain/user"
)

type Schema struct {
	Records []user.Item
}

// Client represents a client to the data store.
type Client struct {
	// Path is the relative filename.
	Path string

	data  *Schema
	mutex sync.RWMutex

	// Services
	userService UserService
}

func NewClient(path string) *Client {
	c := &Client{
		Path: path,
	}
	c.userService.client = c
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

// UserService returns the user service associated with the client.
func (c *Client) UserService() user.Service { return &c.userService }