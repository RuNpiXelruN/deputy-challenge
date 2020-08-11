package neo

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Service type 
type Service interface{
	Driver() neo4j.Driver
	Sess() neo4j.Session
}

// Client type 
type Client struct {
	driver neo4j.Driver
	sess neo4j.Session
}

// NewClient func 
func NewClient() (*Client, error) {
	s, d, err := Connect()
	if err != nil {
		return nil, err
	}

	return &Client{
		driver: d,
		sess: s,
	}, nil
}

// Driver func 
func (c *Client) Driver() neo4j.Driver {
	return c.driver
}

// Sess func 
func (c *Client) Sess() neo4j.Session {
	return c.sess
}

// Connect func 
func Connect() (neo4j.Session, neo4j.Driver, error) {
		// define driver, session and result vars
		// initialize driver to connect to localhost with ID and password
		// driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test", ""))
		// if err != nil {
		// 	fmt.Println("neo4j.NewDriver error", err)
		// }

		if driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test", ""), func(c *neo4j.Config) {
			c.Encrypted = false
		}); err != nil {
			fmt.Println("Error while establishing graph connection")
		}
		// Open a new session with write access
		session, err := driver.Session(neo4j.AccessModeRead)
		if err != nil {
			return nil, nil, err
		}

		return session, driver, nil
}