package pgx

import (
	"context"

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
}

// DBImpl type 
type DBImpl struct {
	conn *pgx.Conn
}

// DB func 
func (d *DBImpl) DB() *pgx.Conn {
	return d.conn
}

// Service type 
type Service interface{}

// Client type 
type Client struct {
	conn *pgx.Conn
}

// NewClient func 
func NewClient(c *pgx.Conn) *Client {
	return &Client{
		conn: c,
	}
}