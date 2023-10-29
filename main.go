package main

import (
	"fmt"
	"github.com/alexeykhan/amocrm"
	"net/http"
	"qr/qr"
	"qr/router"
)

func main() {

	qr.CreateQR()

	router.InitRoutes()
	CreateAmoClient()

	http.ListenAndServe(":8080", nil)

}

func CreateAmoClient() {
	amoCRM := amocrm.New("dc7037a7-7f4e-4c5c-bbad-a93d1b3774d5", "fhieUsfUTyKmlROK7KeumLDQLb6VvCc1c7WI2f8GHo8xmpLhyrgLN1GwPsUR4CCB", "")
	state := amocrm.RandomState()  // store this state as a session identifier
	mode := amocrm.PostMessageMode // options: PostMessageMode, PopupMode

	authURL, err := amoCRM.AuthorizeURL(state, mode)
	if err != nil {
		fmt.Println("failed to get auth url:", err)
		return
	}

	fmt.Println("Redirect user to this URL:")
	fmt.Println(authURL)

	if err := amoCRM.SetDomain("https://onviz-api.ru/amo_deal"); err != nil {
		fmt.Println("set domain:", err)

		return
	}

	token, err := amoCRM.TokenByCode("authorizationCode")
	if err != nil {
		fmt.Println("get token by code:", err)
		return
	}

	fmt.Println("access_token:", token.AccessToken())
	fmt.Println("refresh_token:", token.RefreshToken())
	fmt.Println("token_type:", token.TokenType())
	fmt.Println("expires_at:", token.ExpiresAt().Unix())
}
