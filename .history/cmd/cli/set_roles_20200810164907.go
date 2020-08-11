package cli

import (
	"context"
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy_test/internal/data"
	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pgx"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setRolesCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting roles")
		ctx := context.Background()

		pgConn, err := pgx.NewConn()
		if err != nil {
			log.Printf("pg.NewConn err:", err)
		}
		defer pgConn.Close(ctx)

		rows, err := pgConn.Query(ctx, "select * from deputy.migrations")
		if err != nil {
			fmt.Println("errrr", err)
		}

		pg := pg.NewClient(pgConn)
		neo := neo.NewClient()
		dbClient := db.NewClient(pg, neo)

		err = dbClient.SetRoles()
		if err != nil {
			log.Printf("Error setting roles: %v\n", err)
		} else {
			log.Println("Successfully set roles")
		}
	},
}

func init() {
	rootCmd.AddCommand(setRolesCMD)
}
