package main

import (
	"net/http"
	"qr/qr"
	"qr/router"
)

func main() {

	qr.CreateQR()

	router.InitRoutes()

	//amocrm.GetToken() //1

	//amocrm.CreateComplexDealAndContact() //2

	//amocrm.DealCreate()

	//amocrm.RefreshTokenAuth()

	http.ListenAndServe(":9090", nil)

}
