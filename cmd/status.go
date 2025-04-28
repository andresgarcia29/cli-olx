package cmd

import (
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
		helpers.RunCheckers(false)
	},
}
