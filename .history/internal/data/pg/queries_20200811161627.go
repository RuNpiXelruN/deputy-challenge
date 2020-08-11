package pg

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
	
	_, err = conn.Prepare(ctx, "setUsers", `call setUsers()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	
	_, err = conn.Prepare(ctx, "getSubordinates", `SELECT * FROM getSubordinates($1)`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}

	return nil
}

// Seed func 
func (c *Client) Seed(ctx context.Context) error {
	err := c.SetRoles(ctx)
	if err != nil {
		return err
	}

	err = c.SetUsers(ctx)
	if err != nil {
		return err
	}
	
	return nil
}

// SetRoles func 
func (c *Client) SetRoles(ctx context.Context) error {

	_, err := c.Exec(ctx, "setRoles")
	if err != nil {
		log.Println("Exec setRoles err:", err)
		return err
	}

	return nil
}

// SetUsers func 
func (c *Client) SetUsers(ctx context.Context) error {
	
	_, err := c.Exec(ctx, "setUsers")
	if err != nil {
		log.Println("Exec setUsers err:", err)
		return err
	}

	return nil
}