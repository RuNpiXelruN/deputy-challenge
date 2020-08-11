package pgx

import (
	"context"
	"log"

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

// DB func 
func (d *DBImpl) DB() *pgx.Conn {
	return d.conn
}

// PrepareQueries func 
func (d *DBImpl) PrepareQueries(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Prepare(ctx, "setRoles", `SELECT FROM deputy.setRoles()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	return nil
}

// Service type 
type Service interface{}

// Client type 
type Client struct {
	conn DB
}

// NewClient func 
func NewClient(conn DB) *Client {
	conn.PrepareQueries(context.Background(), conn.DB())
	
	return &Client{
		conn: conn,
	}
}