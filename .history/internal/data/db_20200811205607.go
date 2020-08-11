package db

import (
	"context"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pg"
)

// Client type
type Client struct {
	Pg  pg.Service
	Neo neo.Service
}

// NewClient func
func NewClient(pg pg.Service, neo neo.Service) *Client {
	return &Client{}
}

// WithNeo func 
func (c *Client) WithNeo() *Client {
	return &Client{
		Neo: neo.NewClient(),
	}
}

// WithPostgres func 
func (c *Client) WithPostgres() *Client {
	return &Client{
		Pg: pg.NewClient(),
	}
}

// WithNeoAndPostgres func 
func (c *Client) WithNeoAndPostgres() *Client {
	return &Client{
		Pg: pg.NewClient(),
		Neo: neo.NewClient(),
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

// GetSubordinates func 
func (c *Client) GetSubordinates(ctx context.Context, userID string) ([]User) {

	e := []error{}
	if c.Neo != nil {
		responseBytes, err := c.Neo.GetSubordinates(ctx, userID)
		if err != nil {
			e = append(e, err)
		}
	}

	if c.Pg != nil {
		users, err := c.Pg.GetSubordinates(ctx, userID)
		if err != nil {
			e = append(e, err)
		}
	}
}

// User type
type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
	Role int    `json:"Role"`
}
