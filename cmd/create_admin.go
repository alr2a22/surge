package cmd

import (
	"fmt"
	"surge/api/account/models"
	"surge/internal/db"

	"github.com/spf13/cobra"
)

var (
	createAdminCmd = &cobra.Command{
		Use:   "create-admin",
		Short: "Create admin user command line",
		Long:  `Create admin user command line with username and password`,
		Run: func(cmd *cobra.Command, args []string) {
			createAdmin(cmd)
		},
	}
)

func init() {
	rootCmd.AddCommand(createAdminCmd)
	createAdminCmd.Flags().String("username", "", "admin username")
	createAdminCmd.MarkFlagRequired("username")
	createAdminCmd.Flags().String("password", "", "admin password")
	createAdminCmd.MarkFlagRequired("password")
}

func createAdmin(cmd *cobra.Command) {
	username := cmd.Flag("username").Value.String()
	password := cmd.Flag("password").Value.String()
	DB := db.GetDBConn()
	user := models.User{Username: username, Admin: true}
	user.SetPassword(password)
	err := DB.Create(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("admin user created successfully")
}
