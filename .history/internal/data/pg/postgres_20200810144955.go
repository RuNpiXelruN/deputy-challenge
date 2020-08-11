package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ryboe/q"
)

var (
	host string = "localhost"
	port string = "5432"
	dbName string = "deputychallenge"
	user string = "deputyuser"
	password string = "password"
)

// Service type 
type Service interface{}

// Client type 
type Client struct {
	pool PostgresPool
}

// NewClient func 
func NewClient(cfg Config) *Client {
	return &Client{
		pool: NewPool(cfg),
	}
}

// PostgresPool type 
type PostgresPool interface {
	DB() *pgxpool.Pool
}

// PoolImpl type 
type PoolImpl struct {
	p *pgxpool.Pool
}

// DB func 
func (p *PoolImpl) DB() *pgxpool.Pool {
	return p.p
}

// NewPool func 
func NewPool() *PoolImpl {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", host, port, dbName, user, password))
	if err != nil {
		return nil
	}

	q.Q(pgCfg)

	pool, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		return nil
	}

	return &PoolImpl{
		p: pool,
	}
}