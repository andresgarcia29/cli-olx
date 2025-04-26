package helpers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func IsInternetConnected() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://www.google.com")
	return err == nil
}

func IsConfigPathIsCreated() bool {
	path := os.Getenv("HOME") + "/.olx"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetAllFilesInPath(folder_path string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(folder_path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)

		return nil
	})

	return files, err
}
