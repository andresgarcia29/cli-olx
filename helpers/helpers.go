package helpers

import (
	"net/http"
	"time"
)

func IsInternetConnected() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://www.google.com")
	return err == nil
}
