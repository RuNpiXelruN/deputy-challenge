package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var (
	host string = "127.0.0.1"
	port int = 5432
	dbName string = "deputychallenge"
	user string = "postgres"
	password string = "password"
)

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

// NewConn func 
func NewConn() (*pgx.Conn, error) {
	pgx.ParseConfig()
	pgx.ConnectConfig()
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}
	conn.Config().RuntimeParams = map[string]string{
		"search_path": "deputy",
	}

	return conn, nil
}