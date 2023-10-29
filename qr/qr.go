package qr

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io"
	"log"
	"net/http"
	"os"
	"qr/exolve"
	"qr/exolve/models"
	"strconv"
)

func CreateQR() {
	sendSms := `smsto:+79587329155`
	var size, content string = "256", sendSms
	var codeData []byte
	qrCodeSize, err := strconv.Atoi(size)
	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		log.Println("generate code error: ", err)
	}
	qrCodeFileName := "qrcode.png"
	qrCodeFile, err := os.Create(qrCodeFileName)
	if err != nil {
		log.Println("create file err: ", err)
	}
	defer qrCodeFile.Close()
	qrCodeFile.Write(codeData)
}

func AmoConn(w http.ResponseWriter, r *http.Request) {
	bs, _ := io.ReadAll(r.Body)
	w.Write(bs)
	w.WriteHeader(200)
	jsonData, err := json.Marshal(bs)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	if err != nil {
		log.Println("error marsh sending message: ", err)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	sendSms := `smsto:+79587329155`
	var size, content string = "256", sendSms
	var codeData []byte

	w.Header().Set("Content-Type", "application/json")

	if content == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			"Could not determine the desired QR code content.",
		)
		return
	}

	qrCodeSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Could not determine the desired QR code size.")
		return
	}

	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err),
		)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	// Save the QR code as a .png file
	qrCodeFileName := "qrcode.png"
	qrCodeFile, err := os.Create(qrCodeFileName)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to create QR code file"))
		return
	}
	defer qrCodeFile.Close()
	qrCodeFile.Write(codeData)

	// Now, the QR code is saved as "qrcode.png" in the current working directory

	w.Write(codeData)

	var message models.Message
	if r.Method == "POST" {
		bs, _ := io.ReadAll(r.Body)
		if bs != nil {
			//exolve.GetCount()
			exolve.GetList()
			w.Write(bs)
			w.WriteHeader(200)
			jsonData := json.Unmarshal(bs, &message)
			if jsonData != nil {
				fmt.Println("Message received successfully")
			}
			log.Println("response:", message)
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
