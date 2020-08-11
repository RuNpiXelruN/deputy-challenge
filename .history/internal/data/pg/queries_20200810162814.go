package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func queries(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Prepare(ctx, "setRoles", `SELECT FROM deputy.setRoles()`)
	if err != nil {
		log.Println("conn.Prepare err:", err)
	}
}

// AfterConnetFunc DB Prepared
type AfterConnetFunc func(context.Context, *pgx.Conn) error

// afterConnectForRead creates the read prepared statements that this application uses
func afterConnectForRead(ctx context.Context, conn *pgx.Conn) (err error) {
	_, err = conn.Prepare(ctx, "get_top_segments_metrics", `SELECT * FROM curator.get_top_segments_metrics($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "get_scenario_metrics", `SELECT * FROM curator.get_scenario_metrics($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(ctx, "get_stat_events_metrics", `SELECT * FROM curator.get_stat_events_metrics($1, $2, $3, $4, $5`)
	if err != nil {
		return err
	}

	return nil
}

// afterConnectForWrite creates the write prepared statements that this application uses
func afterConnectForWrite(ctx context.Context, conn *pgx.Conn) (err error) {
	_, err = conn.Prepare(ctx, "upsert_top_segments_metrics", `SELECT FROM curator.upsert_top_segments_metrics($1, $2, $3, $4, $5, $6, $7, $8, $9)`)

	if err != nil {
		return err
	}

	return nil
}
