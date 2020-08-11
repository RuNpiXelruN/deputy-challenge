package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setUsersCMD = &cobra.Command{
	Use:   "setUsers",
	Short: "Sets users to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting users")
	},
}

func init() {
	rootCmd.AddCommand(setUsersCMD)
}