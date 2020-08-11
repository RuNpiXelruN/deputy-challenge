package neo

import (
	"context"
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func (c *Client) prepareStatement(query string, conn bolt.Conn) (bolt.Stmt, error) {

	st, err := conn.PrepareNeo(query)
	if err != nil {
		log.Println("conn.PrepareNeo error", err)
		return nil, err
	}

	return st, nil
}

// Seed func 
func (c *Client) Seed(ctx context.Context) error {

	st, err := c.prepareStatement(seedDB, c.conn)
	if err != nil {
		return err
	}

	_, err = st.ExecNeo(map[string]interface{}{})
	if err != nil {
		log.Println("st.ExecNeo error", err)
		return err
	}

	return nil
}

// GetSubordinates func 
func (c *Client) GetSubordinates(userID string) []User {

	st, err := c.prepareStatement(getSubordinates, c.conn)
	if err != nil {
		return err
	}
}

