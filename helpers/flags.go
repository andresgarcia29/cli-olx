package helpers

import (
	"fmt"
	"os"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/andresgarcia29/cli-uploader/config"
	"github.com/spf13/cobra"
)

func RunCheckers(mute bool) bool {
	error := false
	var status []string

	if IsInternetConnected() {
		status = append(status, "[✅] Connected to the internet")
	} else {
		status = append(status, "[❌] Not connected to the internet")
		error = true
	}

	if auth.IsAuthenticated() {
		status = append(status, "[✅] Authenticated")
	} else {
		status = append(status, "[❌] Not authenticated")
		error = true
	}

	if IsConfigPathIsCreated() {
		status = append(status, "[✅] Init Configuration")
	} else {
		status = append(status, "[❌] Not Init Configuration")
		error = true
	}
	status = append(status, "[📦] Version: "+config.VERSION)

	if !mute {
		for _, s := range status {
			fmt.Println(s)
		}
	}

	return error
}

func PreChecker(cmd *cobra.Command, args []string) {
	if RunCheckers(true) {
		fmt.Println("[❌] Please verify the status of the CLI [check the command 'olx status']")
		os.Exit(1)
	}
}
