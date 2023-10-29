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

	clientID := `c696f967-94ea-4a47-968a-23ee138adf95`
	grantType := `authorization_code`
	clientSecret := `5uhpb6HXKoZVQ1z9gtONKiRTeloINnKxXVbFkisalm51pu1SkhJsxJyLAnxil9Tu`
	code := `def5020089919e0c7466890f7470826cfbdb71a4316958e6af5d98c5ba10628bcaf2dcea3c48f3a71be6181e618a56239529fbe51f90df427ac3fa64a5b41fcbb5ced55c81badb6ed926d5b9d4c034e3026ee27a49b5f83dfbb409c4e8e56e38007462ceca75a07ad340eb976665a33e4a690ab0fe113b2ee10afb0e38830a1f7336972ecfa9bdb24af13c0cf63db9a547dafcc0f9a5127a3122fc3a0cd054387ff2e78f807e4edd0aabf9dc38085cc0049f29c04548a7c157d99eb9724e653e633480302b321bd29e6b7080d80375489a0bcea7e2721890e610165a0e6edfa333870720fc709090999125e9742da8bb88592bcee241198562d9fd2ad079b1ee224176820e0759d0c1d08d5ceabe3f84ad98a9663464c790cf39390d4dfe29846436ba9b263675947daee8531c0a419d74d81c96ece52e43f39122bfa0e0f9355b606c17cf80c6442ce7680098cc139303e490898b2d6e300322e1537807049aaa27a48e27315afce0560c6ddd31170e6703fdf433ee11ebbe525ebd3f5ddf8c61bb6972b2535ad1494f4f617029a92bcf99362beeb4d7578dfcc22985d81f77698a28e0ef840e003d0b11e06a7ee4c48f9b41e8eb90ecafa128a4ab2f7d4d65b96755bf77632a3334e1cd47651850196d0e4aa0671fdca454c57113774445597173bb6230f5189371e9dcf7c0ded0`
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
	//req.Header.Set("Authorization", "Bearer "+AccessToken)
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
