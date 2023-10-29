package amocrm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetDeals(w http.ResponseWriter, r *http.Request) {
	tokenData, err := loadTokenDataFromFile()
	uri := `https://onvizbitrix.amocrm.ru/api/v4/leads`
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("tokenData.AccessToken:::", tokenData.AccessToken)
	req.Header.Set("Authorization", "Bearer "+tokenData.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
	}
	defer resp.Body.Close()
	response, _ := io.ReadAll(resp.Body)
	log.Println("response::", string(response))
	w.Write(response)
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

func loadTokenDataFromFile() (TokenData, error) {
	// Change the filename to the path of your JSON file
	filename := "token_data.json"

	file, err := os.ReadFile(filename)
	if err != nil {
		return TokenData{}, err
	}

	var tokenData TokenData
	err = json.Unmarshal(file, &tokenData)
	if err != nil {
		return TokenData{}, err
	}

	return tokenData, nil
}
