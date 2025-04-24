package cmd

import (
	"fmt"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(debugCmd)
}

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug",
	Long:  "debug",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("! -- Debug -- !")
		user, err := auth.GetUser()
		if err != nil {
			fmt.Println("Error getting user:", err)
			return
		}
		fmt.Println("User:", user.Username)
		fmt.Println("Expiration:", user.Expiration)
	},
}
