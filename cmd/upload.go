package cmd

import (
	"fmt"
	"strings"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/andresgarcia29/cli-uploader/config"
	"github.com/andresgarcia29/cli-uploader/s3"
	"github.com/andresgarcia29/cli-uploader/shortener"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uploadFilePath string
var showUrl bool

func init() {
	cmdUpload.Flags().StringVarP(&uploadFilePath, "path", "p", "", "Destination to upload the file")
	cmdUpload.Flags().BoolVarP(&showUrl, "url", "u", false, "Destination to upload the file")
	cmdUpload.MarkFlagRequired("path")
	rootCmd.AddCommand(cmdUpload)
}

var cmdUpload = &cobra.Command{
	Use:   "upload",
	Short: "Upload files from the cloud",
	Long:  `Upload files from the aws s3 bucket with a signed url`,

	Run: func(cmd *cobra.Command, args []string) {
		service := s3.Service{
			ServerUrl: config.SIGNER_S3_URL,
		}

		fileName := strings.Split(uploadFilePath, "/")[len(strings.Split(uploadFilePath, "/"))-1]
		extensionFile := strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]
		UUIDfilePath := uuid.New().String() + "." + extensionFile

		url, err := service.UploadFile(UUIDfilePath, uploadFilePath, auth.Login())
		if err != nil {
			fmt.Println("[❌] Error to upload the file:", err)
			return
		}

		shortener := shortener.Shortener{
			Url:       url,
			CreatedBy: "olx",
		}
		url, err = shortener.CreateShortUrl()
		if err != nil {
			fmt.Println("[❌] Error to create the short url:", err)
			return
		}

		fmt.Println("[✅] File uploaded successfully")
		fmt.Println("[🔑] Your key to download the file is:", UUIDfilePath)
		if showUrl {
			fmt.Println("[🌍] Your url to download the file is:", url)
		}
	},
}
