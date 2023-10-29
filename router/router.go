package router

import (
	"net/http"
	"qr/amocrm"
	"qr/exolve"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)

	http.HandleFunc("/amo_deal", amocrm.AmoConn)
	http.HandleFunc("/redirect", amocrm.RedirectHandler)

	http.HandleFunc("/send_sms", exolve.SendSms)

}
