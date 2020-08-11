package db

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PGService type 
type PGService interface{}

// PGClient type 
type PGClient struct {}

// NewPGClient func 
func NewPGClient() *PGClient {
	return &PGClient{}
}

// Connect func 
func Connect() {
	pgCfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}
}