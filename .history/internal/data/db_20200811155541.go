package db

import (
	"context"
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

// SeedDatabases func 
func (c *Client) SeedDatabases(ctx context.Context) error {
	err := c.pg.Seed(ctx)
	if err != nil {
		return err
	}

	err = c.neo.Seed(ctx)
	if err != nil {
		return err
	}
	
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