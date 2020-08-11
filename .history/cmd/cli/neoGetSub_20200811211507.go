package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy_test/internal/data"
	"github.com/spf13/cobra"
)

var neoGetSubCMD = &cobra.Command{
	Use:   "neoGetSub",
	Short: "Returns subordinates from the Neo4j database given a userID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(userID) < 1 {
			fmt.Println("You must provide a userID (eg, `--userID 2`)")
			return
		}

		fmt.Printf("\n..fetching subordinates for userID: %v\n", userID)
		ctx := context.Background()

		dbClient := db.NewClient().WithNeo()
		defer dbClient.Neo.Conn().Close()

		users, err := dbClient.GetSubordinates(ctx, userID)
		if err != nil {
			log.Printf("Error fetching subordinates: %v\n", err)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Println("json.Marshal error", err)
			return
		}

		fmt.Printf("Subordinates for userID %v:\n%v\n", userID, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(neoGetSubCMD)
}
