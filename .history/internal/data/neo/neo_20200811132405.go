package neo

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Service type 
type Service interface{}

// Client type 
type Client struct {}

// NewClient func 
func NewClient() *Client {
	return &Client{}
}

// Connect func 
func Connect() bolt.Conn {
	driver := bolt.NewDriver()

	conn, err := driver.OpenNeo("bolt://neo4j:test@localhost:7687")
	if err != nil {
		fmt.Println("driver.OpenNeo error", err)
	}

	return conn
}