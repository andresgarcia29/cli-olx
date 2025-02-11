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

func isTokenValid(tokenString string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if expVal, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(expVal), 0)
			return time.Now().Before(expTime)
		}
	}
	return false
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

	if isTokenValid(authCreds.AccessToken) {
		return true, &authCreds, nil
	}
	return false, &authCreds, nil
}

func IsAuthenticated() bool {
	ok, _, _ := isTokenStillValid()
	return ok
}
