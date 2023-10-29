package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"qr/exolve/models"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)
	http.HandleFunc("/amo_deal", AmoConn)
	http.HandleFunc("/send_sms", SendSms)

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
}
