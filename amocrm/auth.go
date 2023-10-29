package amocrm

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func GetToken() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`
	clientID := "250e6632-c2d2-4986-93a1-e9bd16002327"
	clientSecret := "f6edVVFGIZKEPiAabNTknkALBsqdmTpYAPRD0WEixrA6BLtIA5WMZIQnFZKaT1xT"
	authCode := "def5020007223ff71807070f757754801e4b250513c1f6e9981a010860d94370fc57567ba043d8c88df99e7860a4b80c9cb7a885fa0c006647033adbdf7d93698df1045c8b8f5f5d96d5b2e135da14e8ef7e6311458429f9c7de9681c0d79c4d12bd55f28430878982543ddf2773cead1f4fc6bccb83c7dba1af1feddc0b4b2b9322c81a401b31148dc95556ee0554bf2b138bd595888472f9452da4b8bfea407292267a225b7b0f2575a1015fa35765103a792e63f276a795e0132211e2cab2f71267354e937ce3e62f265efee557239acdda66e8e10ce64f41d3c4e4de553cc5d1a932021d07766a6f13699a1e0774c233afa059cd4618c097961dcc642811124eafe70245ceb90ca933651ba527fefe03aabd698c257d56c863d8f41b70a28646fbd65d3d9867d9af31fce8d72f81561847013fe6caa2d3d8302b8c2155bfce4952d2bb641082abb0fad6d782ff8a5964acbb5945f79fc763e0938b2f2b67b33e977ad7cc3d1b2bd727cabd1baf6c1fc8812acc901e4afd3c77f7c755cb888d3a74e05fdc18733e1b1fbd44195bda228db5dd71f6684f491ba7da2c3e3da0a7063fadd78716f286fe2f54fed44ba53d1a958be9fdc0cf2f59c8f2bcf7e792442833cef319946967111ae96b47905be331d90b3406ea0ba4e77a525387e33e81c9e42c87a2a067852bb8aa630745"
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

	bodyRead, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// 'body' now contains the response data as a byte slice
	// You can convert it to a string if needed
	responseText := string(bodyRead)
	fmt.Println("responseText: ", responseText)
	/*var token Token
	bs, _ := io.ReadAll(r.Body)
	jsonData := json.Unmarshal(bs, &token)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	log.Println("response:", token)*/
}

/*func CreateAmoClient() {
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
}*/

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	var token Token
	bs, _ := io.ReadAll(r.Body)
	jsonData := json.Unmarshal(bs, &token)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	log.Println("response:", token)
}

func AmoConn(w http.ResponseWriter, r *http.Request) {
	bs, _ := io.ReadAll(r.Body)
	w.Write(bs)
	w.WriteHeader(200)
	jsonData, err := json.Marshal(bs)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	if err != nil {
		log.Println("error marsh sending message: ", err)
	}
}
