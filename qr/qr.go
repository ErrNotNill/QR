package qr

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io"
	"log"
	"net/http"
	"qr/exolve"
	"qr/exolve/models"
	"strconv"
)

func HandleRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	sendSms := `smsto:+79587329155`
	var size, content string = request.FormValue("size"), sendSms
	var codeData []byte

	writer.Header().Set("Content-Type", "application/json")

	if content == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired QR code content.",
		)
		return
	}

	qrCodeSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode("Could not determine the desired QR code size.")
		return
	}

	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err),
		)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)

	var message models.Message
	if request.Method == "POST" {

		bs, _ := io.ReadAll(request.Body)
		if bs != nil {

			//exolve.GetCount()
			exolve.GetList()
			writer.Write(bs)
			writer.WriteHeader(200)
			jsonData := json.Unmarshal(bs, &message)
			if jsonData != nil {
				fmt.Println("Message received successfully")
			}
			log.Println("response :", message)
		}
		log.Println("resp:", string(bs))
	}
}

type simpleQRCode struct {
	Content string
	Size    int
}

func (code *simpleQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}
