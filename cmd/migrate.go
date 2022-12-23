package cmd

import (
	"fmt"
	"surge/api/account"
	"surge/api/ride"
	"surge/internal/db"

	"github.com/spf13/cobra"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Database migration",
		Long:  `Database migration`,
		Run: func(cmd *cobra.Command, args []string) {
			migrate(cmd)
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command) {
	DB := db.GetDBConn()
	account.Migrate(DB)
	ride.Migrate(DB)
	fmt.Println("database migrated successfully")
}
