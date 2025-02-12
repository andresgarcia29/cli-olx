package cmd

import (
	"fmt"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/andresgarcia29/cli-uploader/config"
	"github.com/andresgarcia29/cli-uploader/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current version of CLI",
	Long:  `Check the internet connection, version, authentication, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if helpers.IsInternetConnected() {
			fmt.Println("[✅] Connected to the internet")
		} else {
			fmt.Println("[❌] Not connected to the internet")
		}

		if auth.IsAuthenticated() {
			fmt.Println("[✅] Authenticated")
		} else {
			fmt.Println("[❌] Not authenticated")
		}

		if helpers.IsConfigPathIsCreated() {
			fmt.Println("[✅] Init Configuration")
		} else {
			fmt.Println("[❌] Not Init Configuration")
		}

		fmt.Println("[📦] Version: ", config.VERSION)
	},
}
