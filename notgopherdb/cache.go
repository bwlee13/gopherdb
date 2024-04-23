package gopherdb

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Cache struct {
	conn net.Conn
}

// NewCache creates and returns a new Cache instance connected to the specified server address.
func NewCache(serverAddress string) (*Cache, error) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return &Cache{conn: conn}, nil
}

// Set sends a key-value pair to the cache server.
func (c *Cache) Set(key string, value string) error {
	// Prepare the command
	command := fmt.Sprintf("SET %s %s", key, value)
	data := []byte(command)
	// Send the length of the data as a uint32
	if err := binary.Write(c.conn, binary.BigEndian, uint32(len(data))); err != nil {
		return err
	}
	// Send the command data
	_, err := c.conn.Write(data)
	return err
}

// Get retrieves the value for a given key from the cache server.
func (c *Cache) Get(key string) (string, error) {
	command := fmt.Sprintf("GET %s", key)
	data := []byte(command)
	// Send the length of the data as a uint32
	if err := binary.Write(c.conn, binary.BigEndian, uint32(len(data))); err != nil {
		return "", err
	}
	// Send the command data
	if _, err := c.conn.Write(data); err != nil {
		return "", err
	}

	// Read the response (assuming the response will also be prefixed with its length)
	var length uint32
	if err := binary.Read(c.conn, binary.BigEndian, &length); err != nil {
		return "", err
	}
	response := make([]byte, length)
	if _, err := c.conn.Read(response); err != nil {
		return "", err
	}
	return string(response), nil
}

// Close terminates the connection to the cache server.
func (c *Cache) Close() error {
	return c.conn.Close()
}
