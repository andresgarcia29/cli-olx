package auth

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/andresgarcia29/cli-uploader/config"
	"github.com/golang-jwt/jwt"
)

func isLocalAuthentication() (bool, error) {
	_, err := os.Stat(config.OLX_CONFIG_PATH)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type User struct {
	Expiration  int64  `json:"exp"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func GetUser() (*User, error) {
	var user User

	file, err := os.Open(config.OLX_CONFIG_PATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var authCreds AuthenticationCredentials
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(fileContent, &authCreds)

	user.AccessToken = authCreds.AccessToken

	token, _, err := new(jwt.Parser).ParseUnverified(authCreds.AccessToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if expVal, ok := claims["username"].(string); ok {
			user.Username = expVal
		}
		if expVal, ok := claims["exp"].(float64); ok {
			user.Expiration = int64(expVal)
		}
	}
	return &user, nil
}

func isTokenStillValid() (bool, *AuthenticationCredentials, error) {
	file, err := os.Open(config.OLX_CONFIG_PATH)
	if err != nil {
		return false, nil, err
	}
	defer file.Close()

	var authCreds AuthenticationCredentials
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return false, nil, err
	}
	json.Unmarshal(fileContent, &authCreds)

	user, err := GetUser()
	if err != nil {
		return false, nil, err
	}

	expTime := time.Unix(int64(user.Expiration), 0)
	return time.Now().Before(expTime), &authCreds, nil
}

func IsAuthenticated() bool {
	ok, _, _ := isTokenStillValid()
	return ok
}
