package cli

import (
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy_test/internal/data"
	"github.com/runpixelrun/deputy_test/internal/data/neo"
	"github.com/runpixelrun/deputy_test/internal/data/pg"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setRolesCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v\n", rootCmd.Use, "..setting roles")

		pgConn, err := pg.NewConn()
		if err != nil {
			log.Printf("pg.NewConn err:", err)
		}

		pg := pg.NewClient(pgConn)
		neo := neo.NewClient()
		dbClient := db.NewClient(pg, neo)

		err := dbClient.SetRoles()
		if err != nil {
			log.Printf("Error setting roles: %v\n", err)
		} else {
			log.Println("Successfully set roles")
		}
	},
}

func init() {
	rootCmd.AddCommand(setRolesCMD)
}
