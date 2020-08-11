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
var setUsersCMD = &cobra.Command{
	Use:   "setUsers",
	Short: "Sets users to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting users")
		ctx := context.Background()

		pgConn, err := pgx.NewConn()
		if err != nil {
			log.Println("pg.NewConn err:", err)
		}
		defer pgConn.Close(ctx)

		pg := pgx.NewClient(pgx.NewDB())
		neo := neo.NewClient()
		dbClient := db.NewClient(pg, neo)

		err = dbClient.SetUsers(ctx)
		if err != nil {
			log.Printf("Error setting users: %v\n", err)
		} else {
			log.Println("Successfully set users")
		}
	},
}

func init() {
	rootCmd.AddCommand(setUsersCMD)
}