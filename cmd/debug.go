package cmd

import (
	"fmt"

	"github.com/andresgarcia29/cli-uploader/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(debugCmd)
}

var debugCmd = &cobra.Command{
	Use:              "debug",
	Short:            "debug",
	Long:             "debug",
	PersistentPreRun: helpers.PreChecker,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("! -- Debug -- !")
	},
}
