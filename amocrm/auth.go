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

func GetToken() {
	uri := `https://onvizbitrix.amocrm.ru/oauth2/access_token`

	clientID := `1fffc59d-f083-42c7-804e-5e97f954a3a9`
	grantType := `authorization_code`
	clientSecret := `gLhk8pHSg1p8VXK9WdAQRTXFDnaepQ6PNLkVgrlaL0jLs3fEpMzd8HDMZn0bA9i1`
	code := `def502008a42435d716418febdd48f79c7d3ad996277f6913478a37888619f709b00b5f8accd3dbf3af3d2e29c6c2929cecac3abc32e592266599f35cb482461603028c8c90af5266c144ab20f021932badc0b287cc29393b7b8ba8b8ed44a2040de6fbe9c2f99cc9ae8764fe28bf1affa313878ec6c97cb846ddb1dcfd2ee65b8d0bcd2f70d75968e4e2a2e9a7e68bcd2b994f3d6d8534af087b1dedf67808c8f72564f38cdca4284d0405ccc60e6f195e973e9da61d7aa510e7e0e6b68f3f6a82bcd5bd4878d702e37c999d87f3da4bada966e19fa11029dbfb6c31bd3dec6b69ffce613d5ecdfcb2dd469458b859b3f44d2b3185d1385b41f624442c1447bf2610df43069024b2cad9dcfbffd685dbe30c37086da1ba935ce238c9c24ce845b2165a851e82570cc99859732359f2b1b0505da2a27384ddcf5a14a10f0e6456c7d5b1f379324ffd6b5fce2784c008978bd4c62ed9c0875b63bf5c933b6fe31058b846e2c2697ccc7e348bbdbc5cbb6bc889b3a1efc95d599a57bfbc16fe9121fa11b6d0a7a163d9896f08d304cbcf87b1d165cca9b16beaa5759590be4a72c85c5c9ea40df5950e996a7d5b016660bbb30387b7b100c3c43cb43a49d756d5bbb0e51dda6f1c7be9eae4f135e9fdaedb2dd8a26fcfdf2888490f217cd44a126673f2edd89ac4813ea3285`
	redirectUri := `https://onviz-api.ru/amo_deal`

	form := url.Values{}
	form.Add("client_id", clientID)
	form.Add("grant_type", grantType)
	form.Add("client_secret", clientSecret)
	form.Add("code", code)
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

func DealCreate() {

	subdomain := "onvizbitrix"

	leadData := `[
    {
        "name": "Сделка для примера 1",
        "created_by": 0,
        "price": 20000,
        "custom_fields_values": [
            {
                "field_id": 294471,
                "values": [
                    {
                        "value": "Наш первый клиент"
                    }
                ]
            }
        ]
    },
    {
        "name": "Сделка для примера 2",
        "price": 10000,
        "_embedded": {
            "tags": [
                {
                    "id": 2719
                }
            ]
        }
    }
]`

	data := []byte(leadData)
	r := bytes.NewReader(data)
	apiEndpoint := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", subdomain)

	req, err := http.NewRequest("POST", apiEndpoint, r)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Println("ACCESS_TOKEN>>>>>> ", AccessToken)
	req.Header.Set("Authorization", "Bearer "+AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("resp.StatusCode", resp.StatusCode)
	fmt.Println("resp.StatusCode", resp.StatusCode)
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Lead creation failed with status code %d\n", resp.StatusCode)
		return
	}

	fmt.Println("Lead created successfully!")
}

func DealCreateHandler(w http.ResponseWriter, r *http.Request) {

	subdomain := "onvizbitrix"

	leadData := `[
    {
        "name": "Сделка для примера 1",
        "created_by": 0,
        "price": 20000,
        "custom_fields_values": [
            {
                "field_id": 294471,
                "values": [
                    {
                        "value": "Наш первый клиент"
                    }
                ]
            }
        ]
    },
    {
        "name": "Сделка для примера 2",
        "price": 10000,
        "_embedded": {
            "tags": [
                {
                    "id": 2719
                }
            ]
        }
    }
]`

	data := []byte(leadData)
	rs := bytes.NewReader(data)

	apiEndpoint := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", subdomain)

	req, err := http.NewRequest("POST", apiEndpoint, rs)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	fmt.Println("ACCESS_TOKEN>>>>>> ", AccessToken)
	//req.Header.Set("Authorization", "Bearer "+AccessToken)
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
		"pipeline_id": 7407690, // Replace with the actual pipeline ID
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
	req.Header.Set("Authorization", "Bearer "+AccessToken)
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
