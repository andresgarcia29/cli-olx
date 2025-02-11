package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCMD)
}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "Create the initial configuration",
	Long:  `Create the initial configuration for the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		path := os.Getenv("HOME") + "/.olx"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
		}
		fmt.Println("[🔥] All already set up")
	},
}
