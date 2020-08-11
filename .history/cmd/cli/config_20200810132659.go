package cmd

import (
	"bitbucket.org/appcurator/engagement/internal/data/pgx"
)

var cfg config

type config struct {
	Database         pgx.Config
}
