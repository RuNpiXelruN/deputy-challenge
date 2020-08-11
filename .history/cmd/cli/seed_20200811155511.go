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
var seedCMD = &cobra.Command{
	Use:   "seed",
	Short: "Seeds data in both the Postgres and Neo4j databases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..seeding databases")
		ctx := context.Background()

		pgConn, err := pgx.NewConn()
		if err != nil {
			log.Println("pg.NewConn err:", err)
		}
		defer pgConn.Close(ctx)

		pg := pgx.NewClient(pgx.NewDB())
		neo := neo.NewClient()
		dbClient := db.NewClient(pg, neo)

		err = dbClient.SeedDatabases(ctx)
		if err != nil {
			log.Println(err.Error())
			return
		}
		
		log.Println("..successfully seeded databases")
	},
}

func init() {
	rootCmd.AddCommand(seedCMD)
}
