package neo

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Service type 
type Service interface{}

// Client type 
type Client struct {
	Conn bolt.Conn
}

// NewClient func 
func NewClient() *Client {
	
	return &Client{
		Conn: Connect(),
	}
}

// Connect func 
func Connect() bolt.Conn {
		// define driver, session and result vars
		// initialize driver to connect to localhost with ID and password
		driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test", ""))
		if err != nil {
			fmt.Println("neo4j.NewDriver error", err)
		}
		// Open a new session with write access
		session, err := driver.Session(neo4j.AccessModeWrite)
		if err != nil {
			return nil, nil, err
		}
		// return session, driver, nil
	// driver := bolt.NewDriver()
	
	// conn, err := driver.OpenNeo("bolt://neo4j:test@localhost:7687")
	// if err != nil {
	// 	fmt.Println("driver.OpenNeo error", err)
	// }

	return conn
}