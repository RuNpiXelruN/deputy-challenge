package neo

import (
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
func Connect() {
	driver := bolt.NewDriver()

	driver.OpenNeo("bolt://neo4j:test@localhost:7687")
}