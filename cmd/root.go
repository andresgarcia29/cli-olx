package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "olx",
	Short: "olx is a CLI tool to upload files to the cloud",
	Long: `olx is a CLI tool to upload files to the cloud. 
				Also it can download files from the cloud.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, I'm Olx. Happy to help!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
