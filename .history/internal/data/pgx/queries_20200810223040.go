package pgx

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// PrepareQueries func 
func (c *Client) PrepareQueries(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Prepare(ctx, "setRoles", `call setRoles()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	
	_, err := conn.Prepare(ctx, "setUsers", `call setUsers()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	return nil
}