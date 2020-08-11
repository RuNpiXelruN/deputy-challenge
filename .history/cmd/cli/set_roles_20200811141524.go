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
			log.Println("pg.NewConn err:", err)
		}
		defer pgConn.Close(ctx)

		pg := pgx.NewClient(pgx.NewDB())
		neo, err := neo.NewClient()
		if err != nil {
			fmt.Println("neo.NewClient", err)
		}

		defer neo.Driver().Close()
		defer neo.Sess().Close()
		
		dbClient := db.NewClient(pg, neo)

		err = dbClient.SetRoles(ctx)
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
