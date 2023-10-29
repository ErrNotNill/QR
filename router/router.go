package router

import (
	"net/http"
	"qr/amocrm"
	"qr/exolve"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)

	http.HandleFunc("/token", amocrm.GetTokenHandler)
	http.HandleFunc("/deal", amocrm.DealCreateHandler)
	http.HandleFunc("/get_deals", amocrm.GetDeals)

	//http.HandleFunc("/redirect", amocrm.RedirectHandler)

	http.HandleFunc("/send_sms", exolve.SendSms)

}
