package neo

import (
	"context"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func (c *Client) prepareStatement(query string, conn bolt.Conn) bolt.Stmt {

}

// Seed func 
func (c *Client) Seed(ctx context.Context) error {
	return nil
}

