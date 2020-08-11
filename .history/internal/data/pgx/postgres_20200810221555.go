package pgx

import (
	"context"
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

// DB type 
type DB interface {
	PrepareQueries(ctx context.Context, conn *pgx.Conn) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	DB() *pgx.Conn
}

// NewDB func 
func NewDB() *DBImpl {
	conn, err := NewConn()
	if err != nil {
		log.Println("NewConn error:", err)
		os.Exit(1)
	}

	return &DBImpl{
		conn: conn,
	}
}

// Service type 
type Service interface{
	PrepareQueries(ctx context.Context, conn *pgx.Conn) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	DB() *pgx.Conn
}

// Client type 
type Client struct {
	conn *pgx.Conn
}

// NewClient func 
func NewClient(conn DB) *Client {

	err := conn.PrepareQueries(context.Background(), conn.DB())
	if err != nil {
		os.Exit(1)
	}

	return &Client{
		conn: conn,
	}
}

// Exec func 
func (c *Client) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.conn.Exec(ctx, sql, arguments...)
}

// DB func 
func (c *Client) DB() *pgx.Conn {
	return c.conn
}