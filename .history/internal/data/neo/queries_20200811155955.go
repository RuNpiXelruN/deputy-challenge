package neo

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func (c *Client) prepareStatement(query string, conn bolt.Conn) bolt.Stmt {

}