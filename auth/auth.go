package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andresgarcia29/cli-uploader/config"
)

type AuthenticationCredentials struct {
	TokenID      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func Authenticate() (bool, error) {
	isLocalAuthenticated, err := isLocalAuthentication()
	if err != nil {
		fmt.Println(err)
	}

	if isLocalAuthenticated {
		ok, authCreds, err := isTokenStillValid()
		if err != nil {
			log.Fatalln("The token is not longer valid", err)
		}

		if ok {
			fmt.Println("✅ The token is valid")
			return true, nil
		}

		fmt.Println("[❌] The authentication fail, retry...")
		fmt.Println("[🚧] Start refresh token")
		err = getRefreshTokenRequest(authCreds.RefreshToken)
		if err == nil {
			fmt.Println("[🔄] The token was refreshed correctly")
			return true, nil
		}

		fmt.Println("[🚧] Start web authentication")
		isLocalAuthenticated = false
	}

	if !isLocalAuthenticated {
		openURL(getAuthenticationUrl())

		srv := &http.Server{Addr: ":" + config.EPHIMERAL_SERVER_PORT}
		var code string

		http.HandleFunc("/done", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "You are authenticated now! Please close this windows.")
			go getTokenRequest(code, srv)
			fmt.Println("[✅] Authentication completed")
		})

		http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
			code = r.URL.Query()["code"][0]
			http.Redirect(w, r, "http://localhost:"+config.EPHIMERAL_SERVER_PORT+"/done", http.StatusSeeOther)
		})
		srv.ListenAndServe()
		return true, nil
	}

	return false, nil
}

func Login() string {
	fmt.Println("[⏱️] Start authentication...")
	ok, authCreds, err := isTokenStillValid()
	if err != nil {
		log.Fatalln("Error to authenticate", err)
	}
	if !ok {
		fmt.Println("[❌] The authentication fail, retry...")
		err := getRefreshTokenRequest(authCreds.RefreshToken)
		if err != nil {
			fmt.Println(err)
			fmt.Println("[🙏] Please execute <olx auth> to start the authentication process")
			os.Exit(1)
		}
		fmt.Println("[🔄] The token was refreshed correctly")
		return Login()
	}
	fmt.Println("[🔐] The token is valid")
	return authCreds.AccessToken
}
