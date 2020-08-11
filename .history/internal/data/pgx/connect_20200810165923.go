package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// AfterConnetFunc type 
type AfterConnetFunc func(context.Context, *pgx.Conn) error
var xx pgconn.AfterConnectFunc = func(ctx context.Context, pgconn *pgx.PgConn) error {
	return nil	
}

// NewConn func 
func NewConn() (*pgx.Conn, error) {

	// var afterConnect AfterConnetFunc
	// afterConnect = afterConnectForWrite

	conCfg, err := pgx.ParseConfig("postgres://postgres:password@localhost:5432/postgres?search_path=deputy")
	if err != nil {
		fmt.Println("pgx.ParseConfig error", err)
		return nil, err
	}
	conCfg.AfterConnect = xx

	conn, err := pgx.ConnectConfig(context.Background(), conCfg)
	if err != nil {
		fmt.Println("pgx.Connect error", err)
		return nil, err
	}

	return conn, nil
}