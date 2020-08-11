package cli

import (
	"fmt"

	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/ryboe/q"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var neoCMD = &cobra.Command{
	Use:   "neoCMD",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting roles")

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
	rootCmd.AddCommand(neoCMD)
}
