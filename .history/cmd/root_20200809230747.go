package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whiteout",
	Short: "\nThe whiteout service deletes third party user data",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ver string) {
	if err := rootCmd.Execute(); err != nil {
		log.Println("rootCmd.Execute error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&userIDForDelete, "userID", "", "UserID (entityID) of which to delete third party data")
}
