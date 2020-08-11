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
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable search_path=deputy", host, port, dbName, user, password))
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}

	return conn, nil
}

func ConnTest(config configuration.DatabaseConfiguration) (pool *pgx.ConnPool, err error) {
	connectionUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"deputyuser",
		"password",
		"localhost",
		5432,
		"deputychallenge",
		false)

	connectionConfig, err := pgx.ParseURI(connectionUri)

	if err != nil {
		return pool, err
	}

	maxConnections := 50

	poolConfig := pgx.ConnPoolConfig{connectionConfig, maxConnections, nil}

	pool, err = pgx.NewConnPool(poolConfig)

	if err != nil {
		return pool, err
	}

	Database = pool
	return pool, err
}