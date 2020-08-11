package cli

import "github.com/runpixelrun/deputy_test/internal/data/pg"

var cfg config

type config struct {
	PG pg.Config
}
