package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
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

		n := neo.NewClient()
		users, err := n.GetSubordinates(userID)
		if err != nil {
			log.Println("n.GetSubordinates error", err)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Println("json.Marshal error", err)
			return
		}

		fmt.Printf("Subordinates for userID %v -\n%v", userID, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(neoGetSubCMD)
}
