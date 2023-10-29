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
var ClientID string = `a29f8587-7865-4129-8d12-38f2eb4d8479`
var ClientSecret string = `1FbhcKrgVDRjyD6g65TGQ4AOmuCtUNAZFdqLX56ZAjnzh97ys93gk64jK5hPVrs0`
var CodeAuth string = `def50200a2e0a07a28d03e0763c1e00c7911658f2899c771a6432f1395190cf1ecd1399a1cc71c502a25d75089d714f886d29e0637d68a69f73225637fbd3f5244da3fc4ec0a837e52b8b2eff52f2f122165373ea9a80d983a532ec78e0e4a6cd3c617dfcd721f2229c311d1d9e810f337c2667b7ed81734deed2e0632978a622f1ee9e70203423016d59cc8d4e7a3595bee18df7b930c1b4071a1015edf8acd56a483b6224d770c7d321f14e58870a6e74a10af674198d2df1bd695a8339e5d040ecb7cd45b5d037e2b40c8c728dbf771f9f198a7f9f94dc03d8960b8166b484e56021a5c87b1db3ffdcf12dbec26666a16f9b05e7b1a1876a5f3e65cbae909a8a2f091dbba6069182dec9af3c0bb05d1f47e0e29e0a2f444d7f762c2e6bb50c3e2193e9c50abc24e893c2387afaeb85d56bc661064663e25738718ca3c410ff2249a8f9b1ad92f1130fb524f0dbb19cad4a7c6ceb9a0826fae5f26e94cdb89724ccb96f57190ce56b157cbe84d72b771e56ff46397023fcdbe68fcbd407b5277142fb5903f914addb98c8eb511fec1a0eb02259fae69446f05dbdf671be8e4326000ef0a89cec7c1cdff8310a930a29d73438aa14feb0fcc931e7f807da200eba5dfd1170f7ce07aa3cdd954b0b7b51c3613a18213b8d2a541f10eeda904c1abd03cce06db`

func RefreshTokenAuth() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`
	grantType := `refresh_token`
	redirectUri := `http://localhost:9090`

	form := url.Values{}
	form.Add("client_id", ClientID)
	form.Add("grant_type", grantType)
	form.Add("client_secret", ClientSecret)
	form.Add("refresh_token", RefreshToken)
	form.Add("redirect_uri", redirectUri)

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

	//AccessToken = response.AccessToken
	LongRefreshToken = response.RefreshToken

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
	redirectUri := `http://localhost:9090`

	form := url.Values{}
	form.Add("client_id", ClientID)
	form.Add("grant_type", grantType)
	form.Add("client_secret", ClientSecret)
	form.Add("code", CodeAuth)
	form.Add("redirect_uri", redirectUri)

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

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	} else {
		fmt.Println("Request successful!")
	}
}

type CustomFieldValue struct {
	FieldID int `json:"field_id"`
	Values  []struct {
		Value string `json:"value"`
	} `json:"values"`
}

type Embedded struct {
	Tags []struct {
		ID int `json:"id"`
	} `json:"tags"`
}

type Lead struct {
	Name              string             `json:"name"`
	CreatedBy         int                `json:"created_by,omitempty"`
	Price             int                `json:"price"`
	CustomFieldValues []CustomFieldValue `json:"custom_fields_values,omitempty"`
	Embedded          Embedded           `json:"_embedded,omitempty"`
}

func DealCreate() {

	subdomain := "onvizbitrix"
	var jsonStr = []byte(`{"name":"Сделка для примера 1"}`)
	/*leadData := []Lead{
		{
			Name:      "Сделка для примера 1",
			CreatedBy: 0,
			Price:     20000,
			CustomFieldValues: []CustomFieldValue{
				{
					FieldID: 294471,
					Values: []struct {
						Value string `json:"value"`
					}{
						{
							Value: "Наш первый клиент",
						},
					},
				},
			},
		},
		{
			Name:  "Сделка для примера 2",
			Price: 10000,
			Embedded: Embedded{
				Tags: []struct {
					ID int `json:"id"`
				}{
					{
						ID: 2719,
					},
				},
			},
		},
	}*/

	apiEndpoint := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", subdomain)

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	//fmt.Println("ACCESS_TOKEN>>>>>> ", AccessToken)
	req.Header.Set("Authorization", "Bearer "+LongRefreshToken)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("LONG_REFRESH_TOKEN>>>>>> ", LongRefreshToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("resp.StatusCode", resp.StatusCode)
	fmt.Println("resp.StatusCode", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Lead creation failed with status code %d\n", resp.StatusCode)
		return
	}

	fmt.Println("Lead created successfully!")
}

func DealCreateHandler(w http.ResponseWriter, r *http.Request) {
	subdomain := "onvizbitrix"
	leadData := r.Body

	fmt.Println("r.Body: ", r.Body)
	jsonData, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	rs := bytes.NewReader(jsonData)

	apiEndpoint := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", subdomain)

	req, err := http.NewRequest("POST", apiEndpoint, rs)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	//fmt.Println("ACCESS_TOKEN>>>>>> ", AccessToken)
	req.Header.Set("Authorization", "Bearer "+LongRefreshToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Lead creation failed with status code %d\n", resp.StatusCode)
		return
	}

	fmt.Println(r.Body)
	sb, _ := io.ReadAll(resp.Body)
	w.Write(sb)
	fmt.Println("Lead created successfully!")
}

func DealCreateWithPipeLine(w http.ResponseWriter, r *http.Request) {
	subdomain := "onvizbitrix"

	// Define the lead data in a map
	leadData := map[string]interface{}{
		"name":        "NewLeadName",
		"pipeline_id": "7407690", // Replace with the actual pipeline ID
		// Add any other required lead data here
	}

	leadJSON, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling lead data:", err)
		return
	}

	apiEndpoint := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", subdomain)

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(leadJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Println("ACCESS_TOKEN>>>>>> ", AccessToken)
	req.Header.Set("Authorization", "Bearer "+LongRefreshToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Lead creation failed with status code %d\n", resp.StatusCode)
		return
	}
	fmt.Println("Lead created successfully!")
}
