package cli

import (
	"fmt"

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
		n.GetSubordinates(userID)
	},
}

func init() {
	rootCmd.AddCommand(neoGetSubCMD)
}
