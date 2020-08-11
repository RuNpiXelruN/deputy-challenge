package neo

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
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
func (c *Client) GetSubordinates(userID string) ([]User, error) {

	idInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("strconv.Atoi error", err)
		return nil, err
	}

	st, err := c.prepareStatement(getSubordinates, c.conn)
	if err != nil {
		return nil, err
	}

	rows, err := st.QueryNeo(map[string]interface{}{
		"userID": idInt,
	})

	if err != nil {
		log.Println("st.QueryNeo error", err)
		return nil, err
	}

	defer rows.Close()
	data, _, err := rows.All()

	users := c.MapResponseToUsers(data)
	return users, nil
}

// User type 
type User struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	RoleID int `json:"Role"`
}

// MapResponseToUsers func 
func (c *Client) MapResponseToUsers(data [][]interface{}) []User {

	users := []User{}
	
	for _, vals := range data {
		for _, val := range vals {
			u := User{}
			bytes, _ := json.Marshal(val.(graph.Node).Properties)
			json.Unmarshal(bytes, &u)
			users = append(users, u)
		}
	}

	return users
}

// UserResp func 
type UserResp struct {
	Users []User
}

