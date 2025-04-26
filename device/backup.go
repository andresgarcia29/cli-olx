package device

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/andresgarcia29/cli-uploader/helpers"
	"github.com/andresgarcia29/cli-uploader/s3"
)

func GenerateBrewFile(path string) {
	fmt.Println("Starting to create Brewfile...")
	if exec.Command("brew", "bundle", "dump", "--force", "--file="+path).Run() != nil {
		fmt.Println("Error creating Brewfile")
		return
	}
	fmt.Println("Brewfile created successfully")
}

func ExecuteBrewfile(path string) error {
	fmt.Println("Starting to install Brewfile...")
	cmd := exec.Command("brew", "bundle", "install", "--file="+path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error installing Brewfile: %v\n", err)
	}
	fmt.Println(string(output))
	fmt.Println("Brewfile installed successfully")
	return nil
}

func BackUpProcess(deviceFilesList []DeviceFiles, username, accessToken string) {

	s3Service := s3.S3Service{}
	for _, deviceFile := range deviceFilesList {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return
		}
		deviceFile.Path = strings.Replace(deviceFile.Path, "$HOME", homeDir, 1)

		fileName := filepath.Base(deviceFile.Path)
		fileNameUserPath := username + "/device/" + fileName

		if deviceFile.HasIgnoreField {
			content, err := helpers.RemoveIgnoreFileSection(deviceFile.Path)
			if err != nil {
				fmt.Printf("Error removing ignore file section: %v\n", err)
				continue
			}
			path, err := helpers.SaveTempFile(content)
			if err != nil {
				fmt.Printf("Error saving temp file: %v\n", err)
				continue
			}
			deviceFile.Path = path
		}

		if deviceFile.Type == File {
			if deviceFile.GenerateFunction != nil {
				fmt.Println("Generating file...")
				deviceFile.GenerateFunction(deviceFile.Path)
			}
			_, err = s3Service.UploadFile(fileNameUserPath, deviceFile.Path, accessToken)
			if err != nil {
				fmt.Printf("Error uploading file %s: %v\n", deviceFile.Path, err)
				continue
			}
		} else if deviceFile.Type == Folder {
			fmt.Println("Uploading folder...")
			err = s3Service.UploadFolder(username, deviceFile.Path, accessToken)
			if err != nil {
				fmt.Printf("Error uploading folder %s: %v\n", deviceFile.Path, err)
				continue
			}
		}
	}
}
