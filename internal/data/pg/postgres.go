package pg

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	host     string = "127.0.0.1"
	port     int    = 5432
	dbName   string = "deputychallenge"
	user     string = "postgres"
	password string = "password"
)

// Service type
type Service interface {
	PrepareQueries(ctx context.Context, conn *pgx.Conn) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Seed(ctx context.Context) error
	GetSubordinates(ctx context.Context, userID string) ([]byte, error)
	Conn() *pgx.Conn
	DB() *pgx.Conn
}

// Client type
type Client struct {
	conn *pgx.Conn
}

// NewClient func
func NewClient() *Client {
	c := &Client{
		conn: NewDB(),
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

// Conn func
func (c *Client) Conn() *pgx.Conn {
	return c.conn
}

// User type
type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
	Role int    `json:"Role"`
}
