package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setRolesCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v %U\n", rootCmd.Use, "Setting roles...", U+2911)
	},
}

func init() {
	rootCmd.AddCommand(setRolesCMD)
}
