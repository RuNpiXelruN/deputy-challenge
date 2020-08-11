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

	res, err := st.ExecNeo(map[string]string{}{"", ""})
	if err != nil {
		st.ExecNeo
	}

	return nil
}

