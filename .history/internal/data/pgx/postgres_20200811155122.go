package pgx

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	host string = "127.0.0.1"
	port int = 5432
	dbName string = "deputychallenge"
	user string = "postgres"
	password string = "password"
)

// NewDB func 
func NewDB() *pgx.Conn {
	conn, err := NewConn()
	if err != nil {
		log.Println("NewConn error:", err)
		os.Exit(1)
	}

	return conn
}

// Service type 
type Service interface{
	PrepareQueries(ctx context.Context, conn *pgx.Conn) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	DB() *pgx.Conn
	Seed(ctx context.Context) error
}

// Client type 
type Client struct {
	conn *pgx.Conn
}

// NewClient func 
func NewClient(conn *pgx.Conn) *Client {

	c := &Client{
		conn: conn,
	}

	err := c.PrepareQueries(context.Background(), c.conn)
	if err != nil {
		os.Exit(1)
	}

	return c
}

// Exec func 
func (c *Client) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return c.conn.Exec(ctx, sql, args...)
}

// Query func 
func (c *Client) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.conn.Query(ctx, sql, args...)
}

// DB func 
func (c *Client) DB() *pgx.Conn {
	return c.conn
}

// SeedDatabase func 
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

	fmt.Println("setRoles complete")

	return nil
}

// SetUsers func 
func (c *Client) SetUsers(ctx context.Context) error {
	
	_, err := c.Exec(ctx, "setUsers")
	if err != nil {
		log.Println("Exec setUsers err:", err)
		return err
	}

	fmt.Println("setUsers complete")

	return nil
}