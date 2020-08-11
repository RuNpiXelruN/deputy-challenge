package cli

import (
	"context"
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy-challenge/internal/data"
	"github.com/spf13/cobra"
)

var seedCMD = &cobra.Command{
	Use:   "seed",
	Short: "Seeds data in both the Postgres and Neo4j databases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v\n", "..seeding databases")
		ctx := context.Background()

		dbClient := db.NewClient().WithNeoAndPostgres()

		defer dbClient.Pg.Conn().Close(ctx)
		defer dbClient.Neo.Conn().Close()

		err := dbClient.SeedDatabases(ctx)
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println("..complete.")
	},
}

func init() {
	rootCmd.AddCommand(seedCMD)
}
