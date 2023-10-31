package amocrm

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"qr/settings"
)

func GetDeals(w http.ResponseWriter, r *http.Request) {
	//tokenData, err := loadTokenDataFromFile()
	uri := fmt.Sprintf(`https://%s.amocrm.ru/api/v4/leads`, settings.Subdomain)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	//fmt.Println("tokenData.AccessToken:::", tokenData.AccessToken)
	req.Header.Set("Authorization", "Bearer "+settings.AccessToken)
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

func CreateDealAndContact(sender string) {
	uri := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads/complex", settings.Subdomain)

	leadData := fmt.Sprintf(`[
   {
      "name": "Сбор номеров по QR",
      "_embedded":{
         "contacts":[
            {
               "first_name":"QR-Клиент",
              "custom_fields_values":[
                  {
                     "field_id":%s,
                     "values":[
                        {
                           "value":"+%s"}]}]}]}}]`, settings.CrmPhoneField, sender)
	body := []byte(leadData)
	/*jsonData, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))*/

	//rs := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+settings.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	rdr, _ := io.ReadAll(resp.Body)
	fmt.Println("string(rdr): ", string(rdr))
	fmt.Println("Lead created successfully!")
}

func DealCreate() {
	uri := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", settings.Subdomain)
	leadData := fmt.Sprintf(`[
	    {
	        "name": "Сделка для примера 1",
	        "price": 20000,
	    },
	    {
	        "name": "Сделка для примера 2",
	        "price": 10000,
	    }
	]`)
	body := []byte(leadData)
	/*jsonData, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))*/

	//rs := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+settings.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	rdr, _ := io.ReadAll(resp.Body)
	fmt.Println("string(rdr): ", string(rdr))
	fmt.Println("Lead created successfully!")
}

func DealCreateHandler(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads", settings.Subdomain)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	/*leadData := fmt.Sprintf(`[
	    {
	        "name": "Сделка для примера 1",
	        "price": 20000,
	    },
	    {
	        "name": "Сделка для примера 2",
	        "price": 10000,
	    }
	]`)*/

	//body := []byte(leadData)
	/*jsonData, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))*/

	//rs := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+settings.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	rdr, _ := io.ReadAll(resp.Body)
	fmt.Println("string(rdr): ", string(rdr))
	sb, _ := io.ReadAll(resp.Body)
	w.Write(sb)
	fmt.Println("Lead created successfully!")
}

func CreateDealAndContactHandler(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf("https://%s.amocrm.ru/api/v4/leads/complex", settings.Subdomain)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	/*leadData := fmt.Sprintf(`[
	    {
	        "name": "Сделка для примера 1",
	        "price": 20000,
	    },
	    {
	        "name": "Сделка для примера 2",
	        "price": 10000,
	    }
	]`)*/

	//body := []byte(leadData)
	/*jsonData, err := json.Marshal(leadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))*/

	//rs := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+settings.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	rdr, _ := io.ReadAll(resp.Body)
	fmt.Println("string(rdr): ", string(rdr))
	sb, _ := io.ReadAll(resp.Body)
	w.Write(sb)
	fmt.Println("Lead created successfully!")
}
