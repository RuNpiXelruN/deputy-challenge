package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// AfterConnetFunc type 
type AfterConnetFunc func(context.Context, *pgx.Conn) error

// NewConn func 
func NewConn() (*pgx.Conn, error) {

	var afterConnect AfterConnetFunc

	conCfg, err := pgx.ParseConfig("postgres://postgres:password@localhost:5432/postgres?search_path=deputy")
	if err != nil {
		fmt.Println("pgx.ParseConfig error", err)
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), conCfg)
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}

	return conn, nil
}