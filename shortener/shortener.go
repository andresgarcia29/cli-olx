package shortener

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/andresgarcia29/cli-uploader/config"
)

type Shortener struct {
	Url       string `json:"url"`
	CreatedBy string `json:"createdBy"`
}

type ShortenerResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Url     string `json:"url"`
}

func (s *Shortener) CreateShortUrl() (string, error) {
	shortenerData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, config.SHORTENER_URL, bytes.NewBuffer(shortenerData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	shortenerResponse := ShortenerResponse{}
	err = json.NewDecoder(resp.Body).Decode(&shortenerResponse)
	if err != nil {
		return "", err
	}

	return shortenerResponse.Url, nil
}
