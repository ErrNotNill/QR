package amocrm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"qr/amocrm/models"
)

var AccessToken string
var RefreshToken string
var LongRefreshToken string
var ClientID = `db1d8c2f-7572-4ed0-b9d8-5549cf366343`
var ClientSecret = `q4BUIfudr6heIybZpGiXGetqG4c0NiewkJAYnsMeobpsvMudpHoxFg38sSSpfZmp`
var CodeAuth = `def50200733eb5b6838be45719062ab16f52071f205da5fc57531f2456d56df2b22d3b2857e9e8ed614fe1daf40f36896e018475867c74236be0297eb9e96aba9bc9b0d7eadbf66d31fd31573d97630e5b90995f4696e286133eed6de1dd02771805db1973efe962e69610fbeeddadb7a5eb47d3c9e0145fe6cb9c8d2351eee13feeffc89b4ba412335102d990c7f4a5278c4b0f446f9e70b98ce01aa1d25f60e4a983fdfa4ced44638fc48b588c5825e4c3dd1de3c1615bbb1c97dfe6aa05a745949416dc8dd64782539eacb415478f55de5c58ba2e6e2724d1bf48ccf0dd65aebb77f2b5e9ee6d6d322294dcb61d7f0ed28deda9f4c6c7f6763684bd8951330a3c26f07144aafcd17a3717ae2343d0783d57736fbc708fa25b2e542d44c8283d40e823c1ab111e30d820abf0fd4799ec7394be4d2e7ef4e293d2535a1a3c73dd1754b3c46d12f7764eb364f85b793784a6ea59acada727ab0db9c8d12a0c6d62dee6ae09adf7933a3032cd97dc6c14dbc76472b5349ab1eecef2f2fa0c4e7eb259ba65ddedca9833c430469de99c35244376a76f85fb7f578f37ced13a47bdd1594b15b3a417a2fc669d5736d6d77540d8c016ef53c07cec0c4599b883d392f59c3ddc5122d1daf6632e82c3d8f37cbe4807e146e889b140e4fe0b72df585d13c3f653cc`
var RedirectUri = `https://onviz-api.ru`

func RefreshTokenAuth() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	form := url.Values{}
	form.Add("client_id", ClientID)
	form.Add("grant_type", "refresh_token")
	form.Add("client_secret", ClientSecret)
	form.Add("refresh_token", RefreshToken)
	form.Add("redirect_uri", RedirectUri)

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

	// Handle the response here
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

	// Access the fields in the struct
	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	log.Println("LongRefreshToken:", LongRefreshToken)
	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}

func GetToken() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	grantType := `authorization_code`

	form := url.Values{}
	form.Add("client_id", ClientID)
	form.Add("grant_type", grantType)
	form.Add("client_secret", ClientSecret)
	form.Add("code", CodeAuth)
	form.Add("redirect_uri", RedirectUri)

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

	// Handle the response here
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

	// Access the fields in the struct
	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	tokenData := TokenData{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}
	writeTokenDataToFile(tokenData)

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}

func writeTokenDataToFile(data TokenData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling token data:", err)
		return
	}

	// Change the filename to the desired file path
	filename := "token_data.json"

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing token data to file:", err)
	} else {
		fmt.Println("Token data written to file:", filename)
	}
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	grantType := `authorization_code`

	form := url.Values{}
	form.Add("client_id", ClientID)
	form.Add("grant_type", grantType)
	form.Add("client_secret", ClientSecret)
	form.Add("code", CodeAuth)
	form.Add("redirect_uri", RedirectUri)

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

	// Handle the response here
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

	// Access the fields in the struct
	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	tokenData := TokenData{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}
	writeTokenDataToFile(tokenData)

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}

}
