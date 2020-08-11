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
var getSubOrdCMD = &cobra.Command{
	Use:   "getSubordinates",
	Short: "Returns subordinates or passed userID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(userID) < 1 {
			fmt.Println("You must provide a userID (eg, `--userID 2`)")
			return
		}
		
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..fetching subordinates")
		ctx := context.Background()

		pgConn, err := pgx.NewConn()
		if err != nil {
			log.Println("pg.NewConn err:", err)
		}
		defer pgConn.Close(ctx)

		pg := pgx.NewClient(pgx.NewDB())
		neo := neo.NewClient()
		dbClient := db.NewClient(pg, neo)

		err = dbClient.GetSubordinates(ctx, userID)
		if err != nil {
			log.Printf("Error fetching subordinates: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getSubOrdCMD)
}
