package models

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
