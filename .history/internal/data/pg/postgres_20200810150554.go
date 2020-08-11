package pg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ryboe/q"
)

var (
	host string = "localhost"
	port int = 5432
	dbName string = "postgres"
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
func NewClient() *Client {
	conn, err := NewConn()
	if err != nil {
		log.Println("NewConn err:", err)
		os.Exit(1)
	}

	return &Client{
		conn: conn,
	}
}

// Connection type 
type Connection interface {
	DB() *pgx.Conn
}

// PoolImpl type 
type PoolImpl struct {
	 *pgx.Conn
}

// DB func 
func (p *PoolImpl) DB() *pgxpool.Pool {
	return p.
}

// NewPool func 
func NewConn() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", host, port, dbName, user, password))
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}

	q.Q(pool)

	return conn, nil
}