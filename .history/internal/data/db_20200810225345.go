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

// GetSubordinates func 
func (c *Client) GetSubordinates(ctx context.Context, userID string) error {
	intID, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("strconv.Atoi error:", err)
		return err
	}

	_, err := c.pg.Exec(ctx, "getSubordinates", intID)
	if err != nil {
		log.Println("Exec getSubordinates err:", err)
		return err
	}

	fmt.Println("getSubordinates complete")

	return nil
}