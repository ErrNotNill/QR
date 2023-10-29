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

	amocrm.CreateAmoClient()
	amocrm.GetToken()
	amocrm.NewAuth()
	amocrm.GetToken()

	http.ListenAndServe(":9090", nil)

}

/*func withEnv(){
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	login := os.Getenv("AMO_LOGIN")
	key := os.Getenv("AMO_KEY")
	domain := os.Getenv("AMO_DOMAIN")

	api := amocrm.NewAmo(login, key, domain)

	company := api.Lead.Create()
	fmt.Println("errhere")
	company.Name = "TEST"
	id, err := api.Lead.Add(company)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(id)
}*/
