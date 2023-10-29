package router

import (
	"net/http"
	"qr/qr"
)

func InitRoutes() {
	http.HandleFunc("/generate", qr.HandleRequest)
	http.HandleFunc("/amo_deal", qr.AmoConn)
}
