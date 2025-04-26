package cmd

import (
	"fmt"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/andresgarcia29/cli-uploader/device"
	"github.com/spf13/cobra"
)

func init() {
	deviceCmd.AddCommand(deviceBackupCmd)
	deviceCmd.AddCommand(restoreBackupCmd)

	rootCmd.AddCommand(deviceCmd)
}

var DeviceFilesList = []device.DeviceFiles{
	{
		Path:           "$HOME/Library/Application Support/Code/User/settings.json",
		Type:           device.File,
		HasIgnoreField: false,
	},
	{
		Path:           "$HOME/.zshrc",
		Type:           device.File,
		HasIgnoreField: true,
	},
	{
		Path:              "/tmp/Brewfile",
		Type:              device.File,
		HasIgnoreField:    false,
		GenerateFunction:  device.GenerateBrewFile,
		ExecutionFunction: device.ExecuteBrewfile,
	},
}

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "This command is used to configure the device for the user",
	Long:  "This command is used to configure the device for the user [Only Works in MacOS]",
}

var deviceBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup device files",
	Long:  "Backup device files to S3",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Login()
		fmt.Println("! -- [Only Works in MacOS] -- !")

		user, err := auth.GetUser()
		if err != nil {
			fmt.Println("Error getting user:", err)
			return
		}

		device.BackUpProcess(DeviceFilesList, user.Username, user.AccessToken)
		fmt.Println("Device configuration completed successfully.")
	},
}

var restoreBackupCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore device files",
	Long:  "Restore device files from S3",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Login()
		fmt.Println("! -- [Only Works in MacOS] -- !")

		user, err := auth.GetUser()
		if err != nil {
			fmt.Println("Error getting user:", err)
			return
		}

		err = device.RestoreProcess(DeviceFilesList, user.Username, user.AccessToken)
		if err != nil {
			fmt.Println("Error restoring device files:", err)
		}
		fmt.Println("Device configuration completed successfully.")
	},
}
