package cmd

import (
	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with AWS Cognito",
	Long:  `Authenticate with AWS Cognito to get credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Authenticate()
	},
}
