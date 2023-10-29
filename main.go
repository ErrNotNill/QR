package main

import (
	"net/http"
	"qr/amocrm"
	"qr/qr"
	"qr/router"
)

func main() {

	qr.CreateQR()

	router.InitRoutes()

	amocrm.GetToken()
	amocrm.DealCreate()

	http.ListenAndServe(":9090", nil)

}
