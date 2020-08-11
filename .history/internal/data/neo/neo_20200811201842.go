package neo

import (
	"context"
	"encoding/json"
	"log"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Service type 
type Service interface{
	Conn() bolt.Conn
	Seed(ctx context.Context) error
	GetSubordinates(userID string) ([]User, error)
}

// Client type 
type Client struct {
	conn bolt.Conn
}

// NewClient func 
func NewClient() *Client {
	return &Client{
		conn: Connect(),
	}
}

// Conn func 
func (c *Client) Conn() bolt.Conn {
	return c.conn
}

// User type 
type User struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	Role int `json:"Role"`
}

// IncomingUser type 
type IncomingUser struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Role int `json:"role_id"`
}

func (c *Client) mapUserNode(node map[string]interface{}) (User, error) {
	i := IncomingUser{}
	bytes, err := json.Marshal(node)
	if err != nil {
		log.Println("json.Marshal error", err)
		return nil, err
	}

	err = json.Unmarshal(bytes, &i)
	if err != nil {
		log.Println("json.Unmarshal error", err)
		return nil, err
	}

	u := User{
		ID: i.ID,
		Name: i.Name,
		Role: i.Role,
	}

	return u, nil
}