package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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

// connTest func 
func connTest() (pool *pgx.ConnPool, err error) {
	pCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", "localhost", 5432, "deputychallenge", "deputyuser", "password"))
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pCfg)
	if err != nil {
		return nil, err
	}
}