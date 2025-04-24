package cmd

import (
	"fmt"
	"log"

	"github.com/andresgarcia29/cli-uploader/auth"
	"github.com/andresgarcia29/cli-uploader/s3"
	"github.com/spf13/cobra"
)

var downloadS3Key string
var downloadFilePathDestination string

func init() {
	cmdDownload.Flags().StringVarP(&downloadS3Key, "key", "k", "", "Key to download the file from the cloud")
	cmdDownload.Flags().StringVarP(&downloadFilePathDestination, "path", "p", "~/.olx/", "Destination to save the file")
	cmdDownload.MarkFlagRequired("key")
	cmdDownload.MarkFlagRequired("path")
	rootCmd.AddCommand(cmdDownload)
}

func addKeyToThePath(path, key string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	path += key
	return path
}

var cmdDownload = &cobra.Command{
	Use:   "download",
	Short: "Download files from the cloud",
	Long:  `Download files from the aws s3 bucket with a signed url`,

	Run: func(cmd *cobra.Command, args []string) {
		s3SignerService := s3.S3SignerService{}

		downloadFilePathDestination = addKeyToThePath(downloadFilePathDestination, downloadS3Key)

		err := s3SignerService.DownloadFile(downloadS3Key, downloadFilePathDestination, auth.Login())
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("[📂] File downloaded successfully")
	},
}
