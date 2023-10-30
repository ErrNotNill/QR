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
var ClientID = `5f47ad30-d132-443d-81fd-db5ab0f05f4c`
var ClientSecret = `wtHTsEmRyYGeJdIkj20wLFOAtsvVYdnhgEefEURhsysp4Ew5xrDu8amy8WCaBk7k`
var CodeAuth = `def50200cae674d90568edcd5d469e0c3bdc495e76e9fa0ece3b65b534b2c0363f155c0072f500cb6e9a835406a975289764536ece54a03f86571bf16d164362fd16cb315bb69348aba45ef2721d861bbe49c838e29901532396493c225376bc62ad3ad671109a94cd68c4477fa830030b5e4cffca95daa201feb1dd5a6c69281d9dd0173f351e6ab14ee6c8ebbaf0e7dbf48a81cc8dde2cd64b317c4463d755047b190ca2758bf39a733d9790b456c6fc66a4623e16ebe48031463005f30b4ac7a615d14f7bdcc973875a9abf7bc419f46a44936af93d4fd58a8d90257ef9b39954dcfbae2f47e465571616c80904154f9acb9b63af880e3246ee33f11380401d73a0c65910e1d0639cc1b5387a5cff6a67eafb42989df299022539e9fa80cae9dd97d8914ea2e9f55596d2bc75656f2c1da1dab5d0c039b3307bd8f54ff38308284c8eb4f2ed5e6f5ea10c7c7dc77c3a5716a65bda3b8864ce5bb1d656ce88ae563696311f96c60c557adcde07f506600a5cd45e4e3fa03b1b47a53c4327f6a251cf4550053c06c845e242043e56ce84a6709105b0e85a24cb5811673e9cf205c748ed426bc6c8c0d92f1d836741a85862c79554898092b89b61b341729530cff56abfab26b825968b818947d89497c4451a3fc3fb20662716e3548a289dffc18304a8ce`
var RedirectUri = `https://onviz-api.ru`

func RefreshTokenAuthHandler(w http.ResponseWriter, r *http.Request) {
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

	reqExample := fmt.Sprintf(`{
  "client_id": "%s",
  "client_secret": "%s",
  "grant_type": "authorization_code",
  "code": "%s",
  "redirect_uri": "%s"
}`, ClientID, ClientSecret, CodeAuth, RedirectUri)

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

	// Handle the response here
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

	// Access the fields in the struct
	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	/*tokenData := TokenData{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}
	writeTokenDataToFile(tokenData)*/

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

	reqExample := fmt.Sprintf(`{
  "client_id": "%s",
  "client_secret": "%s",
  "grant_type": "authorization_code",
  "code": "%s",
  "redirect_uri": "%s"
}`, ClientID, ClientSecret, CodeAuth, RedirectUri)

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

	// Handle the response here
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

	// Access the fields in the struct
	fmt.Println("Token Type:", response.TokenType)
	fmt.Println("Expires In:", response.ExpiresIn)
	fmt.Println("Access Token:", response.AccessToken)
	fmt.Println("Refresh Token:", response.RefreshToken)

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	w.Write([]byte("AccessToken: " + AccessToken))
	w.Write([]byte("RefreshToken: " + RefreshToken))
	/*tokenData := TokenData{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}
	writeTokenDataToFile(tokenData)*/

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}

}
