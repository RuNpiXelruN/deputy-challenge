package neo_test

import (
	"context"
	"testing"

	db "github.com/runpixelrun/deputy-challenge/internal/data"
	"github.com/runpixelrun/deputy-challenge/internal/data/neo"
	"github.com/runpixelrun/deputy-challenge/internal/data/pg"
	"github.com/stretchr/testify/assert"
)

func TestNeo(t *testing.T) {
	ctx := context.Background()

	// [x] Seed should be called

	t.Run("Seed should be called", func(t *testing.T) {

		pg := &pg.ServiceMock{
			SeedFunc: func(ctx context.Context) error {
				return nil
			},
		}
		neo := &neo.ServiceMock{
			SeedFunc: func(ctx context.Context) error {
				return nil
			},
		}

		dbClient := &db.Client{
			Pg:  pg,
			Neo: neo,
		}
		dbClient.SeedDatabases(ctx)

		assert.Len(t, neo.SeedCalls(), 1)
	})
}
