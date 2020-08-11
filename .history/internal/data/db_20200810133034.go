package db

import (
	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pg"
)

// Client type
type Client struct {
	pg pg.Service
	neo neo.Service
}

// NewClient func 
func NewClient(pg pg.Service, neo neo.Service) *Client {
	return &Client{
		pg: pg,
		neo: neo,
	}
}

// SetRoles func 
func (c *Client) SetRoles() error {
	return nil
}