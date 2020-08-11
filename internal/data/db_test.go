package db

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pg"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	ctx := context.Background()

	t.Run("SeedDatabases", func(t *testing.T) {

		// [x] Pg.Seed should be called
		// [x] Neo.Seed should be called
		// [x] if seeds succeed error should nil

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
		dbClient := &Client{
			Pg: pg,
			Neo: neo,
		}

		err := dbClient.SeedDatabases(ctx)

		t.Run("Pg.Seed should be called", func(t *testing.T) {
			assert.Len(t, pg.SeedCalls(), 1)
		})
		
		t.Run("Neo.Seed should be called", func(t *testing.T) {
			assert.Len(t, neo.SeedCalls(), 1)
		})
		
		t.Run("if seeds succeed error should nil", func(t *testing.T) {
			assert.Nil(t, err)
		})
	})

	t.Run("GetSubordinates", func(t *testing.T) {

		// [x] if client.Neo is not nil, Neo.GetSubordinates should be called
		// [x] if client.Pg is not nil, Pg.GetSubordinates should be called
		// [x] if Neo.GetSubordinates succeeds, a []User should be returned and nil error
		// [x] if Pg.GetSubordinates succeeds, a []User should be returned and nil error

		t.Run("if client.Neo is not nil, Neo.GetSubordinates should be called", func(t *testing.T) {
			neo := &neo.ServiceMock{
				GetSubordinatesFunc: func(ctx context.Context, userID string) ([]byte, error) {
					return nil, errors.New("expected error")
				},
			}
			c := &Client{
				Neo: neo,
			}

			c.GetSubordinates(ctx, "23")
			assert.Len(t, neo.GetSubordinatesCalls(), 1)
		})
		
		t.Run("if client.Pg is not nil, Pg.GetSubordinates should be called", func(t *testing.T) {
			pg := &pg.ServiceMock{
				GetSubordinatesFunc: func(ctx context.Context, userID string) ([]byte, error) {
					return nil, errors.New("expected error")
				},
			}
			c := &Client{
				Pg: pg,
			}

			c.GetSubordinates(ctx, "23")
			assert.Len(t, pg.GetSubordinatesCalls(), 1)
		})

		t.Run("if Neo.GetSubordinates succeeds, a []User should be returned and nil error", func(t *testing.T) {
			neo := &neo.ServiceMock{
				GetSubordinatesFunc: func(ctx context.Context, userID string) ([]byte, error) {
					users := []neo.User{
						{ID: 04, Name: "Sawyer", Role: 5},
						{ID: 21, Name: "Brooks", Role: 2},
					}
					bytes, _ := json.Marshal(users)
					return bytes, nil
				},
			}
			c := &Client{
				Neo: neo,
			}

			users, err := c.GetSubordinates(ctx, "23")
			assert.Nil(t, err)
			assert.IsType(t, []User{}, users)
		})
		
		t.Run("if Pg.GetSubordinates succeeds, a []User should be returned and nil error", func(t *testing.T) {
			pg := &pg.ServiceMock{
				GetSubordinatesFunc: func(ctx context.Context, userID string) ([]byte, error) {
					users := []pg.User{
						{ID: 04, Name: "Sawyer", Role: 5},
						{ID: 21, Name: "Brooks", Role: 2},
					}
					bytes, _ := json.Marshal(users)
					return bytes, nil
				},
			}
			c := &Client{
				Pg: pg,
			}

			users, err := c.GetSubordinates(ctx, "23")
			assert.Nil(t, err)
			assert.IsType(t, []User{}, users)
		})
	})
}