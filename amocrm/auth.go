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

func NewAuth() {
	subdomain := "onvizbitrix"
	link := "https://" + subdomain + ".amocrm.ru/oauth2/access_token"

	data := map[string]string{
		"client_id":     "250e6632-c2d2-4986-93a1-e9bd16002327",
		"client_secret": "f6edVVFGIZKEPiAabNTknkALBsqdmTpYAPRD0WEixrA6BLtIA5WMZIQnFZKaT1xT",
		"grant_type":    "authorization_code",
		"code":          "def5020007223ff71807070f757754801e4b250513c1f6e9981a010860d94370fc57567ba043d8c88df99e7860a4b80c9cb7a885fa0c006647033adbdf7d93698df1045c8b8f5f5d96d5b2e135da14e8ef7e6311458429f9c7de9681c0d79c4d12bd55f28430878982543ddf2773cead1f4fc6bccb83c7dba1af1feddc0b4b2b9322c81a401b31148dc95556ee0554bf2b138bd595888472f9452da4b8bfea407292267a225b7b0f2575a1015fa35765103a792e63f276a795e0132211e2cab2f71267354e937ce3e62f265efee557239acdda66e8e10ce64f41d3c4e4de553cc5d1a932021d07766a6f13699a1e0774c233afa059cd4618c097961dcc642811124eafe70245ceb90ca933651ba527fefe03aabd698c257d56c863d8f41b70a28646fbd65d3d9867d9af31fce8d72f81561847013fe6caa2d3d8302b8c2155bfce4952d2bb641082abb0fad6d782ff8a5964acbb5945f79fc763e0938b2f2b67b33e977ad7cc3d1b2bd727cabd1baf6c1fc8812acc901e4afd3c77f7c755cb888d3a74e05fdc18733e1b1fbd44195bda228db5dd71f6684f491ba7da2c3e3da0a7063fadd78716f286fe2f54fed44ba53d1a958be9fdc0cf2f59c8f2bcf7e792442833cef319946967111ae96b47905be331d90b3406ea0ba4e77a525387e33e81c9e42c87a2a067852bb8aa630745",
		"redirect_uri":  "https://onviz-api.ru/redirect",
	}

	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", link, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "amoCRM-oAuth-client/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 204 {
		fmt.Println("Error:", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	access_token, _ := response["access_token"].(string)
	refresh_token, _ := response["refresh_token"].(string)
	token_type, _ := response["token_type"].(string)
	expires_in, _ := response["expires_in"].(float64)

	fmt.Println("Access Token:", access_token)
	fmt.Println("Refresh Token:", refresh_token)
	fmt.Println("Token Type:", token_type)
	fmt.Println("Expires In:", expires_in)
}

func GetToken() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	data := map[string]string{
		"client_id":     "250e6632-c2d2-4986-93a1-e9bd16002327",
		"client_secret": "f6edVVFGIZKEPiAabNTknkALBsqdmTpYAPRD0WEixrA6BLtIA5WMZIQnFZKaT1xT",
		"grant_type":    "authorization_code",
		"code":          "def5020007223ff71807070f757754801e4b250513c1f6e9981a010860d94370fc57567ba043d8c88df99e7860a4b80c9cb7a885fa0c006647033adbdf7d93698df1045c8b8f5f5d96d5b2e135da14e8ef7e6311458429f9c7de9681c0d79c4d12bd55f28430878982543ddf2773cead1f4fc6bccb83c7dba1af1feddc0b4b2b9322c81a401b31148dc95556ee0554bf2b138bd595888472f9452da4b8bfea407292267a225b7b0f2575a1015fa35765103a792e63f276a795e0132211e2cab2f71267354e937ce3e62f265efee557239acdda66e8e10ce64f41d3c4e4de553cc5d1a932021d07766a6f13699a1e0774c233afa059cd4618c097961dcc642811124eafe70245ceb90ca933651ba527fefe03aabd698c257d56c863d8f41b70a28646fbd65d3d9867d9af31fce8d72f81561847013fe6caa2d3d8302b8c2155bfce4952d2bb641082abb0fad6d782ff8a5964acbb5945f79fc763e0938b2f2b67b33e977ad7cc3d1b2bd727cabd1baf6c1fc8812acc901e4afd3c77f7c755cb888d3a74e05fdc18733e1b1fbd44195bda228db5dd71f6684f491ba7da2c3e3da0a7063fadd78716f286fe2f54fed44ba53d1a958be9fdc0cf2f59c8f2bcf7e792442833cef319946967111ae96b47905be331d90b3406ea0ba4e77a525387e33e81c9e42c87a2a067852bb8aa630745",
		"redirect_uri":  "https://onviz-api.ru/redirect",
	}

	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(payload))
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

func CreateAmoClient() {

	clientID := "250e6632-c2d2-4986-93a1-e9bd16002327"
	clientSecret := "f6edVVFGIZKEPiAabNTknkALBsqdmTpYAPRD0WEixrA6BLtIA5WMZIQnFZKaT1xT"
	//grantType := "grant_type":    "authorization_code",
	//authCode :=  "def5020007223ff71807070f757754801e4b250513c1f6e9981a010860d94370fc57567ba043d8c88df99e7860a4b80c9cb7a885fa0c006647033adbdf7d93698df1045c8b8f5f5d96d5b2e135da14e8ef7e6311458429f9c7de9681c0d79c4d12bd55f28430878982543ddf2773cead1f4fc6bccb83c7dba1af1feddc0b4b2b9322c81a401b31148dc95556ee0554bf2b138bd595888472f9452da4b8bfea407292267a225b7b0f2575a1015fa35765103a792e63f276a795e0132211e2cab2f71267354e937ce3e62f265efee557239acdda66e8e10ce64f41d3c4e4de553cc5d1a932021d07766a6f13699a1e0774c233afa059cd4618c097961dcc642811124eafe70245ceb90ca933651ba527fefe03aabd698c257d56c863d8f41b70a28646fbd65d3d9867d9af31fce8d72f81561847013fe6caa2d3d8302b8c2155bfce4952d2bb641082abb0fad6d782ff8a5964acbb5945f79fc763e0938b2f2b67b33e977ad7cc3d1b2bd727cabd1baf6c1fc8812acc901e4afd3c77f7c755cb888d3a74e05fdc18733e1b1fbd44195bda228db5dd71f6684f491ba7da2c3e3da0a7063fadd78716f286fe2f54fed44ba53d1a958be9fdc0cf2f59c8f2bcf7e792442833cef319946967111ae96b47905be331d90b3406ea0ba4e77a525387e33e81c9e42c87a2a067852bb8aa630745"
	redirectUri := "https://onviz-api.ru/amo_deal"

	accessToken := "def50200836cea7fdf4571649c804979a65dd5b0776fc8201c8c66120670303e5b052505e98ea22e5d5ac9b75881fffef780cd1175257ee049afc3132e8b1faca7ea95d31657546ea50b878238b31bd6593bc4ded65d79a11fe5d637472d820766a40547d24f789c83fca6c1da65b06635a21cfa93256630d02cf8cdc0e9896415f5e1f7143fd1a34a76d97600f32de3ab95086a85a58d3dee54451601f37b52d44bd5ca8762906804831f4a94f666b662dcca60bcbfc682f12965d24947d3a0c8708c9fed28c69008193a1bbafb28c301df33033b397a21dbb581ab9cbd982394146e25b2b906ebb13d7cade17e371ba5191c4ba1f3f40cee65d6e0f88f65765c8113fcfa057e763bcc5a0212eff36a05d98e56907d5de1d1b97a6ff791e7787fee41ce0416f1055526f562bc9cc2da9ad7a4091623289bc48ed4ff801818980d795c9de4628a8c0f3e8d5fe9b9fb0d5a7dcf16b25e3c9ce7c5fead6802d30ab664230f02345043bb27ee6d8d6e1496dcb4eb9ad40f9904f8178f1962bbcfccebd2a769f40bfcfff211293a8e079a36d8dd3116d65896e6ba00dfae86927902923ef64ed871ac75c7ca794bceddac26d8f5c424839d945e2ff672866fd27da10d22fd7d8c5cb16cd002fcb99e956cbbb0f8b7e7f7d5f79b9f08a777a169a329a6da1c7759e71dd1488f1099e4cf21"
	amoCRM := amocrm.New(clientID, clientSecret, redirectUri)
	//state := amocrm.RandomState()  // store this state as a session identifier
	mode := amocrm.PostMessageMode // options: PostMessageMode, PopupMode
	authURL, err := amoCRM.AuthorizeURL("8d09edf332cf73a06bb3b6ca95cbf1d0", mode)
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
