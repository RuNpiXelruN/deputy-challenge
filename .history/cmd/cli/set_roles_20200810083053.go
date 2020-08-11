package cli

import (
	"fmt"

	"github.com/runpixelrun/deputy_test/internal/db"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setRolesCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting roles")

		pg := db.NewPGClient()
		neo := db.NewNeoClient()
		client := db.NewClient(pg, neo)
	},
}

func init() {
	rootCmd.AddCommand(setRolesCMD)
}
