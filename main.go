package main

import (
	"encoding/json"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/generate", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	// Replace this content with your desired content for the QR code
	content := "Your QR Code Content Here"

	// Save a copy of the original "TEST.png" file
	err := copyFile("TEST.png", "TEST_original.png")
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(fmt.Sprintf("Error copying the original file: %v", err))
		return
	}

	// Generate and save the QR code as "TEST.png"
	err = generateQRCode(content, "TEST.png")
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(fmt.Sprintf("Error generating QR code: %v", err))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("QR code generated and saved as TEST.png")
}

func copyFile(src, dst string) error {
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, source, 0644)
	if err != nil {
		return err
	}
	return nil
}

func generateQRCode(content, filename string) error {
	qrCodeSize := 256 // Set your desired QR code size
	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}

	codeData, err := qrCode.Generate()
	if err != nil {
		return err
	}

	// Save the QR code as the specified filename
	err = os.WriteFile(filename, codeData, 0644)
	if err != nil {
		return err
	}

	return nil
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

func FindFile(substr string) string {
	var substring string
	file, err := os.ReadFile("settings")
	if err != nil {
		fmt.Println("Error reading")
	}
	list := strings.Split(string(file), ",")
	for k, _ := range list {
		withoutSpaces := strings.Join(strings.Fields(list[k]), "")
		if strings.Contains(withoutSpaces, substr) {
			pattern := `<([^>]+)>`
			re := regexp.MustCompile(pattern)
			match := re.FindString(withoutSpaces)
			if match != "" {
				// Remove the "<" and ">" symbols
				substring = match[1 : len(match)-1]
			} else {
				fmt.Println("No match found")
			}
		}
	}
	return substring
}
