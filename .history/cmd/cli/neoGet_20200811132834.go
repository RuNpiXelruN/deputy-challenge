package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var neoCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting roles")
		
		n := neo.New
	},
}

func init() {
	rootCmd.AddCommand(neoCMD)
}
