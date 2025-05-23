package s3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/andresgarcia29/cli-uploader/config"
	"github.com/andresgarcia29/cli-uploader/helpers"
)

type S3Service struct {
	ServerUrl string
}

func (s *S3Service) getSignUrl(key, operation, access_token string) (*S3ResponseSignUrl, error) {

	S3RequestSignUrl := S3RequestSignUrl{
		Key: key,
	}
	payload, err := json.Marshal(S3RequestSignUrl)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.SIGNER_S3_URL+"/"+operation, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", access_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		log.Println(resp.Status, ": Failed to request")
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("failed to request download url" + string(bodyBytes))
	}

	S3ResponseSignUrl := S3ResponseSignUrl{}
	err = json.NewDecoder(resp.Body).Decode(&S3ResponseSignUrl)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to receive download url")
	}
	if err != nil {
		return nil, err
	}

	return &S3ResponseSignUrl, nil
}

func (s *S3Service) DownloadFile(key, filePath, authorization_token string) error {
	downloadSignUrl, err := s.getSignUrl(key, config.DOWNLOAD_OPERATION, authorization_token)
	if err != nil {
		return err
	}

	resp, err := http.Get(downloadSignUrl.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		log.Println(resp.Status, ": Failed to request download url")
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return errors.New("file not exists")
		}
		return errors.New("failed to download file")
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Service) UploadFolder(usernameID, folderPath, authorization_token string) error {
	filesToUpload, err := helpers.GetAllFilesInPath(folderPath)
	if err != nil {
		return err
	}

	folderName := filepath.Base(folderPath)
	fmt.Println("**")
	fmt.Println(folderPath)
	fmt.Println(folderName)
	fmt.Println(filesToUpload)
	fmt.Println("**")
	for _, file := range filesToUpload {
		fmt.Println("x")
		fmt.Println(file)
		key := usernameID + "/device/" + folderName + "/" + file

		uploadSignUrl, err := s.getSignUrl(key, config.UPLOAD_OPERATION, authorization_token)
		if err != nil {
			return err
		}

		file, err := os.Open(file)
		if err != nil {
			return err
		}
		defer file.Close()

		buffer := bytes.NewBuffer(nil)
		if _, err := io.Copy(buffer, file); err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPut, uploadSignUrl.Url, buffer)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "multipart/form-data")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return errors.New("failed to upload file: " + string(bodyBytes))
		}
	}

	return nil
}

func (s *S3Service) UploadFile(key, filePath, authorization_token string) (string, error) {
	uploadSignUrl, err := s.getSignUrl(key, config.UPLOAD_OPERATION, authorization_token)
	if err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, uploadSignUrl.Url, buffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", errors.New("failed to upload file: " + string(bodyBytes))
	}

	urlToDownload, err := s.getSignUrl(key, config.DOWNLOAD_OPERATION, authorization_token)
	if err != nil {
		return "", err
	}

	return urlToDownload.Url, nil
}
