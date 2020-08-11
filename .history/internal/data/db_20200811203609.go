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
	return &Client{
		Pg:  pg,
		Neo: neo,
	}
}

// WithNeo func 
func (c *Client) WithNeo() *Client {
	return &Client{
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
