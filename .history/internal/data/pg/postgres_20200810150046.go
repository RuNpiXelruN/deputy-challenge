package pg

import (
	"context"
	"fmt"

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
	pool PostgresPool
}

// NewClient func 
func NewClient() *Client {
	return &Client{
		pool: NewPool(),
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
	pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password))

	pool, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		fmt.Println("pgxpool.ConnectConfig err:", err)
		return nil
	}

	q.Q(pool)

	return &PoolImpl{
		p: pool,
	}
}