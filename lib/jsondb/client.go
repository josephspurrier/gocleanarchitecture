package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

// Schema represents the database structure.
type Schema struct {
	Records []interface{}
}

// Client represents a client to the data store.
type Client struct {
	// Path is the relative filename.
	Path string

	data  *Schema
	mutex sync.RWMutex
}

// New takes a folder path and returns a new database client. All json files
// are stored in a folder.
func New(folderPath string) *Client {
	c := &Client{
		Path: folderPath,
		data: new(Schema),
	}

	return c
}

// Reads opens/initializes the database.
func (c *Client) read(recordType string) error {
	var err error
	var b []byte

	// Set the file path.
	filePath := path.Join(c.Path, recordType+".json")

	c.mutex.Lock()

	// Check if the file exists.
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
		if err != nil {
			c.mutex.Unlock()
			return err
		}
	}

	// Read the file.
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	// Read the data into the struct.
	c.data = new(Schema)
	err = json.Unmarshal(b, &c.data)

	c.mutex.Unlock()

	return err
}

// Write saves the database.
func (c *Client) write(recordType string) error {
	var err error
	var b []byte

	// Set the file path.
	filePath := path.Join(c.Path, recordType+".json")

	c.mutex.Lock()

	// Convert the struct to bytes.
	b, err = json.Marshal(c.data)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	// Write the bytes to a file.
	err = ioutil.WriteFile(filePath, b, 0644)
	if err != nil {
		c.mutex.Unlock()
		return err
	}

	c.mutex.Unlock()

	return err
}

// AddRecord adds a record to the database.
func (c *Client) AddRecord(recordType string, rec interface{}) error {
	// Load the data.
	err := c.read(recordType)
	if err != nil {
		return err
	}

	c.data.Records = append(c.data.Records, rec)

	// Save the record to the database.
	return c.write(recordType)
}

// Records retrieves all records from the database.
func (c *Client) Records(recordType string) ([]interface{}, error) {
	// Load the data.
	err := c.read(recordType)
	return c.data.Records, err
}
