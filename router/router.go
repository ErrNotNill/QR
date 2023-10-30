package router

import (
	"net/http"
	"qr/amocrm"
	"qr/exolve"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)
	http.HandleFunc("/send_sms", exolve.SendSms)

	http.HandleFunc("/deal", amocrm.DealCreateHandler)
	http.HandleFunc("/deal_contact", amocrm.CreateDealAndContactHandler)
	http.HandleFunc("/get_deals", amocrm.GetDeals)

	//this routes for testing requests
	//http.HandleFunc("/redirect", amocrm.RedirectHandler)
	//http.HandleFunc("/token", amocrm.GetTokenHandler)
	//http.HandleFunc("/refresh_token", amocrm.RefreshTokenAuthHandler)
}
