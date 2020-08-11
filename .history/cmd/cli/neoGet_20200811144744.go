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

		neo, err := neo.NewClient()
		if err != nil {
			fmt.Println("neo.NewClient", err)
		}

		defer neo.Driver().Close()
		defer neo.Sess().Close()

		cypher := `
			MATCH (u:User {id: 3}) return u
		`

		result, err := neo.Sess().Run(cypher, nil)
		if err != nil {
			fmt.Println("Sess().Run error", err)
		}

		if result.Next(){
			q.Q(result.Record().GetByIndex(0))
		}
	},
}

func init() {
	rootCmd.AddCommand(neoCMD)
}
