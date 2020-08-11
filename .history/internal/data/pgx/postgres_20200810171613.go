package pgx

import (
	"context"
	"log"
	"os"

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
	DB() *pgx.Conn
}

// DBImpl type 
type DBImpl struct {
	conn *pgx.Conn
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

// DB func 
func (d *DBImpl) DB() *pgx.Conn {
	return d.conn
}

// Service type 
type Service interface{}

// Client type 
type Client struct {
	conn DB
}

// NewClient func 
func NewClient(conn DB) *Client {

	err := conn.PrepareQueries(context.Background(), conn.DB())
	if err != nil {
		log.Println("conn.PrepareQueries error:", err)
		os.Exit(1)
	}

	return &Client{
		conn: conn,
	}
}