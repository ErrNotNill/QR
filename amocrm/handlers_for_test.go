package amocrm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"qr/amocrm/models"
	"qr/settings"
)

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	reqExample := fmt.Sprintf(`{
  "client_id": "%s",
  "client_secret": "%s",
  "grant_type": "authorization_code",
  "code": "%s",
  "redirect_uri": "%s"
}`, settings.ClientID, settings.ClientSecret, settings.CodeAuth, settings.RedirectUri)

	body := []byte(reqExample)

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bodyRead, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var response models.Token
	err = json.Unmarshal(bodyRead, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	settings.AccessToken = response.AccessToken
	settings.RefreshToken = response.RefreshToken

	w.Write([]byte("AccessToken: " + settings.AccessToken))
	w.Write([]byte("RefreshToken: " + settings.RefreshToken))

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}

func RefreshTokenAuthHandler(w http.ResponseWriter, r *http.Request) {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	form := url.Values{}
	form.Add("client_id", settings.ClientID)
	form.Add("grant_type", "refresh_token")
	form.Add("client_secret", settings.ClientSecret)
	form.Add("refresh_token", settings.RefreshToken)
	form.Add("redirect_uri", settings.RedirectUri)

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(form.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var response models.Token
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	settings.AccessToken = response.AccessToken
	settings.RefreshToken = response.RefreshToken

	log.Println("LongRefreshToken:", settings.LongRefreshToken)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}
