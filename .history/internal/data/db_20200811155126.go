package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pgx"
)

// Client type
type Client struct {
	pg pgx.Service
	neo neo.Service
}

// NewClient func 
func NewClient(pg pgx.Service, neo neo.Service) *Client {
	return &Client{
		pg: pg,
		neo: neo,
	}
}

// SeedDatabase func 
func (c *Client) SeedDatabases(ctx context.Context) error {
	c.
}

// SetRoles func 
func (c *Client) SetRoles(ctx context.Context) error {

	c.pg.SeedDatabase()
	
	_, err := c.pg.Exec(ctx, "setRoles")
	if err != nil {
		log.Println("Exec setRoles err:", err)
		return err
	}

	// c.neo.SeedDB

	fmt.Println("setRoles complete")

	return nil
}

// SetUsers func 
func (c *Client) SetUsers(ctx context.Context) error {
	
	_, err := c.pg.Exec(ctx, "setUsers")
	if err != nil {
		log.Println("Exec setUsers err:", err)
		return err
	}

	fmt.Println("setUsers complete")

	return nil
}

// User type 
type User struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	Role int `json:"Role"` 
}

// GetSubordinates func 
func (c *Client) GetSubordinates(ctx context.Context, userID string) ([]User, error) {
	intID, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("strconv.Atoi error:", err)
		return nil, err
	}

	rows, err := c.pg.Query(ctx, "getSubordinates", intID)
	if err != nil {
		log.Println("Query getSubordinates err:", err)
		return nil, err
	}
	defer rows.Close()

	result := []User{}
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Role)
		if err != nil {
			log.Println("rows.Scan error", err)
			return nil, err
		}
		result = append(result, u)
	}
	if len(result) < 1 {
		return nil, nil
	}

	return result, nil

}