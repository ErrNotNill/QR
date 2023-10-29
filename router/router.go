package router

import (
	"net/http"
	"qr/amocrm"
	"qr/exolve"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)
	http.HandleFunc("/deal", amocrm.DealCreateHandler)
	http.HandleFunc("/deal_pipe", amocrm.DealCreateWithPipeLine)

	//http.HandleFunc("/redirect", amocrm.RedirectHandler)

	http.HandleFunc("/send_sms", exolve.SendSms)

}
