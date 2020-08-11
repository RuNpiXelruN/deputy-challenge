package pgx

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// AfterConnectFunc func 
type AfterConnectFunc func(ctx context.Context, conn *pgx.Conn) error

// PrepareQueries func 
func (c *Client) PrepareQueries(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Prepare(ctx, "setRoles", `SELECT FROM deputy.setRoles()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	return nil
}
