package cli

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var userID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DeputyTest\n",
	Short: "\nDeputy tech challenge by Justin Davidson (justindavidson23@gmail.com)",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("rootCmd.Execute error:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&userID, "userID", "", "UserID of which to find sub-ordinates of")
}
