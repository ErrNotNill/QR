package amocrm

import (
	"encoding/json"
	"fmt"
	"os"
	"qr/amocrm/models"
)

func writeTokenDataToFile(data models.TokenData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling token data:", err)
		return
	}
	// Change the filename to the desired file path
	filename := "token_data.json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing token data to file:", err)
	} else {
		fmt.Println("Token data written to file:", filename)
	}
}

func loadTokenDataFromFile() (models.TokenData, error) {
	// Change the filename to the path of your JSON file
	filename := "token_data.json"
	file, err := os.ReadFile(filename)
	if err != nil {
		return models.TokenData{}, err
	}
	var tokenData models.TokenData
	err = json.Unmarshal(file, &tokenData)
	if err != nil {
		return models.TokenData{}, err
	}
	return tokenData, nil
}
