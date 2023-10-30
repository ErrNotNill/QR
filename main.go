package main

import (
	"log"
	"net/http"
	"qr/router"
)

func main() {

	//phoneField := settings.FindCustomFields("Чат")
	//webhookUrl := settings.FindSettings("DatabaseUrl")

	//qr.CreateQR()

	router.InitRoutes()
	//amocrm.GetToken()
	//amocrm.CreateDealAndContact()

	//amocrm.DealCreate()

	//amocrm.RefreshTokenAuth()

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println(err.Error())
		return
	}
}
