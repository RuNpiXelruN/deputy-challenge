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

// SetRoles func 
func (c *Client) SetRoles(ctx context.Context) error {
	
	_, err := c.pg.Exec(ctx, "setRoles")
	if err != nil {
		log.Println("Exec setRoles err:", err)
		return err
	}

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
func (c *Client) GetSubordinates(ctx context.Context, userID string) error {
	intID, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("strconv.Atoi error:", err)
		return err
	}

	rows, err := c.pg.Query(ctx, "getSubordinates", intID)
	if err != nil {
		log.Println("Query getSubordinates err:", err)
		return err
	}
	defer rows.Close()

	result := []User{}
	for rows.Next() {
		u := User{}
		rows.Scan(&u.ID)
	}


	fmt.Println("getSubordinates complete")

	return nil
}