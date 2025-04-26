package device

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andresgarcia29/cli-uploader/helpers"
	"github.com/andresgarcia29/cli-uploader/s3"
)

func RestoreProcess(deviceFilesList []DeviceFiles, username, accessToken string) error {
	s3Service := s3.S3Service{}
	for _, deviceFile := range deviceFilesList {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return err
		}
		deviceFile.Path = strings.Replace(deviceFile.Path, "$HOME", homeDir, 1)

		fileName := filepath.Base(deviceFile.Path)
		fileNameUserPath := username + "/device/" + fileName

		if deviceFile.Type == File {
			deviceFilePath := deviceFile.Path
			if deviceFile.HasIgnoreField {
				deviceFilePath = "/tmp/" + fileName
			}
			err = s3Service.DownloadFile(fileNameUserPath, deviceFilePath, accessToken)
			if err != nil {
				fmt.Printf("Error downloading file %s: %v\n", deviceFilePath, err)
				return err
			}
			if deviceFile.HasIgnoreField {
				currentIgnore, err := helpers.GetIgnoreFileSection(deviceFile.Path)
				if err != nil {
					fmt.Printf("Error getting ignore file section: %v\n", err)
					return err
				}
				temporalFile, err := helpers.ReadFileAsLines(deviceFilePath)
				if err != nil {
					fmt.Printf("Error reading file: %v\n", err)
					return err
				}
				joinedFiles := strings.Join(temporalFile, "\n") + currentIgnore
				err = os.WriteFile(deviceFile.Path, []byte(joinedFiles), 0644)
				if err != nil {
					fmt.Printf("Error writing ignore file section: %v\n", err)
					return err
				}
			}
			fmt.Println("[✅] File downloaded successfully:", deviceFile.Path)

			if deviceFile.ExecutionFunction != nil {
				fmt.Println("Execute file installation...")
				err := deviceFile.ExecutionFunction(deviceFile.Path)
				if err != nil {
					fmt.Printf("Error execution file: %v\n", err)
					return err
				}
			}
		}
	}
	return nil
}
