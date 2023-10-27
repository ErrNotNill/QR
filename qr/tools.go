package qr

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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
