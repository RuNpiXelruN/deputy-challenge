package pg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// NewConn func
func NewConn() (*pgx.Conn, error) {
	conCfg, err := pgx.ParseConfig("postgres://postgres:password@127.0.0.1:5433/depchallenge")
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

// NewDB func
func NewDB() *pgx.Conn {
	conn, err := NewConn()
	if err != nil {
		log.Println("NewConn error:", err)
		os.Exit(1)
	}

	return conn
}
