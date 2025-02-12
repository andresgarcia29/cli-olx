package helpers

import (
	"net/http"
	"os"
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
