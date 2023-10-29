package amocrm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alexeykhan/amocrm"
	"io"
	"log"
	"net/http"
)

type Token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetToken(r *http.Request) {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`
	clientID := "dc7037a7-7f4e-4c5c-bbad-a93d1b3774d5"
	clientSecret := "fhieUsfUTyKmlROK7KeumLDQLb6VvCc1c7WI2f8GHo8xmpLhyrgLN1GwPsUR4CCB"
	authCode := "def50200afc33242d5f17bb037ff864ab9cd9cf9f318a052647f3d3b22e63e58cf53a6f3ff58ec54b6610dcc7c9378b159a7a56e347ac39afe83a3c33db0a98fa18c9f0ec5eefc7aa1a5babe431f9643ed930982cfaf16fc1fbeba6d27ffaf0e43881c7c9bb1f174b3ddbc5593c0a827fde632d73924f97fbde77ddeca1e08f3d0211f64dc308f70904ea5007a78e62d39f38d82a78ab9be3bf61e6080a21f16d1e08e2720ab079ec11b6242e2e9d22bc2c53339c1af71dd3b9a998fea72618edaefe5a597b45d9824e2b671551afb145dbf3f4acfd96cf9baedd759b664f8aac96b5da9c38a1f7b7044039952e5ffb3d153bc334a829d440bdc68272935e5455294491f48b63bb18dcc0868fbd43ec1707966c2bdb005ad9cd9703534ccb8ceaf411d5932caa903004625450ed680127b308ff1adbc19d9fcdc505a83c3f0e7651b6423ff3393683723acc176b0039f0039ac9cc3e84b4f730159b23ac7ceac05f903b589f4f5456ed6adaa4bfdcb624537650258f8e7ad0513bef23237388f4ffe61b0d8fba2d23d1b26da33241c838d89132f6cb40839c6a7de61157ae33224a32f6973767dd2e36251fe0cfcabbf2b3960b00c0176d39a289f79020d25b894d949caade8b725452070539385c5708d785b97840abdf54f11700be9f3c17b0556bd7ac02b8429072b9c60e71cb0"
	redirectUri := "https://onviz-api.ru/redirect"

	bodyStr := fmt.Sprintf(`{
  "client_id": "%s",
  "client_secret": "%s",
  "grant_type": "authorization_code",
  "code": "%s",
  "redirect_uri": "%s"
}`, clientID, clientSecret, authCode, redirectUri)
	body := []byte(bodyStr)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		log.Println("error creating request")
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var token Token
	bs, _ := io.ReadAll(r.Body)
	jsonData := json.Unmarshal(bs, &token)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	log.Println("response:", token)
}

func CreateAmoClient() {
	accessToken := "def50200afc33242d5f17bb037ff864ab9cd9cf9f318a052647f3d3b22e63e58cf53a6f3ff58ec54b6610dcc7c9378b159a7a56e347ac39afe83a3c33db0a98fa18c9f0ec5eefc7aa1a5babe431f9643ed930982cfaf16fc1fbeba6d27ffaf0e43881c7c9bb1f174b3ddbc5593c0a827fde632d73924f97fbde77ddeca1e08f3d0211f64dc308f70904ea5007a78e62d39f38d82a78ab9be3bf61e6080a21f16d1e08e2720ab079ec11b6242e2e9d22bc2c53339c1af71dd3b9a998fea72618edaefe5a597b45d9824e2b671551afb145dbf3f4acfd96cf9baedd759b664f8aac96b5da9c38a1f7b7044039952e5ffb3d153bc334a829d440bdc68272935e5455294491f48b63bb18dcc0868fbd43ec1707966c2bdb005ad9cd9703534ccb8ceaf411d5932caa903004625450ed680127b308ff1adbc19d9fcdc505a83c3f0e7651b6423ff3393683723acc176b0039f0039ac9cc3e84b4f730159b23ac7ceac05f903b589f4f5456ed6adaa4bfdcb624537650258f8e7ad0513bef23237388f4ffe61b0d8fba2d23d1b26da33241c838d89132f6cb40839c6a7de61157ae33224a32f6973767dd2e36251fe0cfcabbf2b3960b00c0176d39a289f79020d25b894d949caade8b725452070539385c5708d785b97840abdf54f11700be9f3c17b0556bd7ac02b8429072b9c60e71cb0"
	amoCRM := amocrm.New("dc7037a7-7f4e-4c5c-bbad-a93d1b3774d5", "fhieUsfUTyKmlROK7KeumLDQLb6VvCc1c7WI2f8GHo8xmpLhyrgLN1GwPsUR4CCB", "https://onviz-api.ru/amo_deal")
	state := amocrm.RandomState()  // store this state as a session identifier
	mode := amocrm.PostMessageMode // options: PostMessageMode, PopupMode
	authURL, err := amoCRM.AuthorizeURL(state, mode)
	if err != nil {
		fmt.Println("failed to get auth url:", err)
		return
	}
	fmt.Println("Redirect user to this URL:")
	fmt.Println(authURL)
	if err := amoCRM.SetDomain("https://onvizbitrix.amocrm.ru"); err != nil {
		fmt.Println("set domain:", err)
		return
	}
	token, err := amoCRM.TokenByCode(accessToken)
	if err != nil {
		fmt.Println("get token by code:", err)
		return
	}
	fmt.Println("access_token:", token.AccessToken())
	fmt.Println("refresh_token:", token.RefreshToken())
	fmt.Println("token_type:", token.TokenType())
	fmt.Println("expires_at:", token.ExpiresAt().Unix())
}
