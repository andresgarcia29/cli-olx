package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/andresgarcia29/cli-uploader/config"
)

func getAuthenticationUrl() string {
	return `https://us-east-1bbn6lmi7e.auth.us-east-1.amazoncognito.com/login/continue?client_id=6e75vs6eqq7ghql61o8cm21ig5&response_type=code&redirect_uri=http://localhost:8888/code&scope=email+openid+phone`
}

func getTokenRequest(code string, srv *http.Server) error {
	cognitoUrl := "https://us-east-1bbn6lmi7e.auth.us-east-1.amazoncognito.com/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", config.AuthenticationClientConfig.ClientId)
	data.Set("client_secret", config.AuthenticationClientConfig.ClientSecret)
	data.Set("redirect_uri", config.AuthenticationClientConfig.RedirectUrl)
	data.Set("code_verifier", "GENERATED_VERIFIER")

	u, err := url.ParseRequestURI(cognitoUrl)
	if err != nil {
		return err
	}
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	file, err := os.Create(config.OLX_CONFIG_PATH)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}

	srv.Shutdown(context.Background())
	return nil
}

func getRefreshTokenRequest(refresh_token string) error {
	cognitoUrl := "https://us-east-1bbn6lmi7e.auth.us-east-1.amazoncognito.com/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refresh_token)
	data.Set("client_id", config.AuthenticationClientConfig.ClientId)
	data.Set("client_secret", config.AuthenticationClientConfig.ClientSecret)
	data.Set("redirect_uri", config.AuthenticationClientConfig.RedirectUrl)
	data.Set("code_verifier", "GENERATED_VERIFIER")

	u, err := url.ParseRequestURI(cognitoUrl)
	if err != nil {
		return err
	}
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var authCreds AuthenticationCredentials
	err = json.NewDecoder(res.Body).Decode(&authCreds)
	if err != nil {
		return err
	}
	authCreds.RefreshToken = refresh_token

	authCredsData, err := json.Marshal(authCreds)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("[ℹ️] Error to refresh token body: ", res.Body)
		return errors.New("[❌] error to refresh token")
	}

	file, err := os.Create(config.OLX_CONFIG_PATH)
	if err != nil {
		return err
	}
	file.Write(authCredsData)

	return nil
}
