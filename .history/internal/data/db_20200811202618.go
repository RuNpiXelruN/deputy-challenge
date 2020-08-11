package db

import (
	"context"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pg"
)

// Client type
type Client struct {
	pg  pg.Service
	neo neo.Service
}

// NewClient func
func NewClient(pg pg.Service, neo neo.Service) *Client {
	return &Client{
		pg:  pg,
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
