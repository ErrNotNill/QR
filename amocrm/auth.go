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
)

var AccessToken string
var RefreshToken string
var LongRefreshToken string
var ClientID = `fb311eff-bf34-452d-8134-8fecb67d4529`
var ClientSecret = `KzYHv76cDySmZjFBpdKpHwRYhXCVqX4tsYKhj1VEZv5exCijeV5RA934Ns051NUR`
var CodeAuth = `def50200367892ca284029240a6db37fbb38d239801426dabdeb792ba4e2ab74e794736318ee5e3a9bdbc3f38164c7285ecbe05b66fbe9b438d28deba9a95aa48081dbf7477c63c8720cd7b5ecfc1aff89688e14413a6bf8945343bdf1053c3600c4200a0116fc7f6bbaad05728cfc397251429ba88b0c159546b86d1e8e63d9de0480c31760f423606af7aded081a31a721dcb767c4410498217857830d01f329af439a92038dd0a2edaf01077c5a044896d22e1824a3347afc21ea8c8cb9f1dbebea30bcd0174a7775549639f2007bc9792f175f56cba1610dd052da5a0e3a9012aed5b65bd82bc7dcaac547cd793e1a4ef40ea340e9b7b9f6041fbbee434b7ceed26fbfe4eb20903a3e68d35284eae9e7405dc2064217289821d1e7e8965bdfec3cec4cf224766b3af1859f056f3688971f704f0518d687d2a1b52b7ea5367df63febef41b428a165f8b6b5994d1ed183a0879ace349f72f75d9e264377052aea5856e7d4275ec632f5494817d8c3c1034fbe383700335051c5ce7f3980123ceca4a44a5df5f0aaef19d340d472e8ed5ec09ed2a3571916d32f9f40777eba11bb1532f7e1dd21cc408e00514eb570ff9c8a18c5ea616e31e5f6cc126daf4290e83b57603dc43bb82434deef342e5e3644133c0b66386b8083427ba964c6e69679603564`
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

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	log.Println("LongRefreshToken:", LongRefreshToken)

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

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}
