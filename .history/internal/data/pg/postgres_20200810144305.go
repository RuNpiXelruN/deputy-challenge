package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Config type 
type Config struct {
	User                  string
	Password              string
	Host                  string
	Port                  int
	DBName                  string
}

// Service type 
type Service interface{}

// Client type 
type Client struct {
	pool Pool
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
func NewPool(cfg Config) *PoolImpl {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password))
	if err != nil {
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		return nil
	}

	return &PoolImpl{
		p: pool,
	}
}