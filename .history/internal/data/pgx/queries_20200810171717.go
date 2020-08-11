package pgx

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// PrepareQueries func 
func (d *DBImpl) PrepareQueries(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Prepare(ctx, "setRoles", `SELECT FROM setRoles()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
		return err
	}
	return nil
}