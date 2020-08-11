package cli

import (
	"fmt"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/ryboe/q"
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
		cypher := `
			MATCH (u:User) return u
		`

		rows, err := n.Conn().QueryNeo(cypher, nil)
		if err != nil {
			fmt.Println("n.Conn.ExecNeo error", err)
		}

		defer rows.Close()

		data, _, err := rows.All()
		if err != nil {
			fmt.Println("rows.All error", err)
		}
		
		q.Q(data)
	},
}

func init() {
	rootCmd.AddCommand(neoGetSubCMD)
}
