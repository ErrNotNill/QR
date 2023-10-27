package main

import (
	"net/http"
	"qr/router"
)

func main() {

	router.InitRoutes()

	http.ListenAndServe(":8080", nil)

}
