package neo

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Service type 
type Service interface{
	Conn() bolt.Conn
}

// Client type 
type Client struct {
	conn bolt.Conn
}

// NewClient func 
func NewClient() *Client {

	return &Client{
		conn: Con2(),
	}
}

// Conn func 
func (c *Client) Conn() bolt.Conn {
	return c.conn
}

// Connect func 
func Connect() bolt.Conn {
	driver := bolt.NewDriver()

	neoConnString := fmt.Sprintf("bolt://%s:%s@%s:7687", "neo4j", "test", "localhost")

	conn, err := driver.OpenNeo(neoConnString)
	if err != nil {
		fmt.Printf("Error getting neo connection: %+v\n", err)
	}

	fmt.Println("\n\n\nconn", conn)
	return conn
}