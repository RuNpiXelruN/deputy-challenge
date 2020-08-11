package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy-challenge/internal/data"
	"github.com/spf13/cobra"
)

var neoGetSubCMD = &cobra.Command{
	Use:   "neoGetSub",
	Short: "Returns subordinates from the Neo4j database given a userID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(userID) < 1 {
			fmt.Println("You must provide a userID (eg, `--userID 3`)")
			return
		}

		ctx := context.Background()

		dbClient := db.NewClient().WithNeo()
		defer dbClient.Neo.Conn().Close()

		users, err := dbClient.GetSubordinates(ctx, userID)
		if err != nil {
			log.Printf("Error fetching subordinates: %v\n", err)
			return
		}

		if len(users) < 1 {
			fmt.Printf("No subordinates found for userID: %v\n", userID)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Println("json.Marshal error:", err)
			return
		}

		fmt.Printf("\nSubordinates for userID %v (from Neo4j):\n%v\n", userID, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(neoGetSubCMD)
}
