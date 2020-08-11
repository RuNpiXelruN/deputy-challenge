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
type Service interface{
	Pool() *pgxpool.Pool
}

// Client type 
type Client struct {
	pool Pool
}

// NewClient func 
func NewClient(p Pool) *Client {
	return &Client{
		pool: p,
	}
}

// Pool func 
func (c *Client) Pool() *pgxpool.Pool {
	return c.pool
}

// Pool type 
type Pool interface {}


// Connect func 
func Connect(cfg Config) (*pgxpool.Pool, error) {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		return nil, err
	}

	return pool, nil
}