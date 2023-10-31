package exolve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"qr/amocrm"
	"qr/exolve/models"
	"qr/settings"
	"strconv"
	"time"
)

func SendSms(w http.ResponseWriter, r *http.Request) {
	var message models.IncomingMessage
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading message")
	}
	err = json.Unmarshal(bs, &message)
	if err != nil {
		log.Println("error unmarshaling message: ", err)
	}
	w.Write([]byte(message.Sender)) //sender number phone
	fmt.Fprint(w, message)
	log.Println("message.Sender:", message.Sender)
	log.Println("message:", message)

	amocrm.CreateDealAndContact(message.Sender) //there we create deal and contact in AmoCrm when we get sms
}

func CreateBody(datestart string) []byte {
	dateGte, err := time.Parse(time.RFC3339, datestart)
	if err != nil {
		fmt.Println(err)
	}
	dateLte := time.Now().Format(time.RFC3339)

	command := fmt.Sprintf(`{
      "date_gte": "%s",
      "date_lte": "%s"
    }`, dateGte.Format(time.RFC3339), dateLte)
	return []byte(command)
}

func GetList() {
	uri := `https://api.exolve.ru/messaging/v1/GetList`

	dateStart := "2023-10-20T15:04:05.000Z" //YYYY-MM-DD
	//dateFinish := "2023-10-27T15:04:05.000Z" //YYYY-MM-DD
	//dateFinish we don't use, because instead dateFinish now we got current date
	body := CreateBody(dateStart)

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		log.Println("")
	}
	req.Header.Add("Authorization", "Bearer "+settings.ExolveAccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, _ := io.ReadAll(resp.Body)

	var response models.MessageResponse
	err = json.Unmarshal(bs, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	/*for _, message := range response.Messages {
		fmt.Printf("Message ID: %s\n", message.MessageID)
		fmt.Printf("Application UUID: %s\n", message.ApplicationUUID)
		fmt.Printf("Date: %s\n", message.Date.Format(time.RFC3339))
		fmt.Printf("Number: %s\n", message.Number)
		fmt.Printf("Sender: %s\n", message.Sender)
		fmt.Printf("Receiver: %s\n", message.Receiver)
		fmt.Printf("Text: %s\n", message.Text)
		fmt.Printf("Direction: %d\n", message.Direction)
		fmt.Printf("Segments Count: %d\n", message.SegmentsCount)
		fmt.Printf("Billing Status: %d\n", message.BillingStatus)
		fmt.Printf("Delivery Status: %d\n", message.DeliveryStatus)
		fmt.Printf("Channel: %d\n", message.Channel)
		fmt.Printf("Status: %d\n", message.Status)
	}*/

}

func GetCount() {
	uri := `https://api.exolve.ru/messaging/v1/GetCount`

	dateStart := "2023-10-20T15:04:05.000Z" //YYYY-MM-DD
	//dateFinish := "2023-10-27T15:04:05.000Z" //YYYY-MM-DD
	body := CreateBody(dateStart)

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		log.Println("")
	}
	req.Header.Add("Authorization", "Bearer "+settings.ExolveAccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var response models.Count
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error:", err)
		return
	}
	count, _ := strconv.Atoi(response.Count)
	fmt.Println("Count:", count)

}
