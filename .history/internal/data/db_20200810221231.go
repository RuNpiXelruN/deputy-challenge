package db

import (
	"context"
	"log"

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
	c.pg.Ex
	// _, err := c.pg.DB().Exec(ctx, "setRoles")
	if err != nil {
	// 	log.Println("Exec setRoles err:", err)
	// 	return err
	// }

	return nil
}