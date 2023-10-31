package main

import (
	"log"
	"net/http"
	"qr/amocrm"
	"qr/qr"
	"qr/router"
	"qr/settings"
)

func main() {

	qr.CreateQR()
	settings.InitSettings() //in this package we customize settings, example ClientID, ClientSecret and other...
	router.InitRoutes()     //there we have more additional methods
	amocrm.GetToken()

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println(err.Error())
		return
	}

}
