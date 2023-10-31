package settings

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var AccessToken string
var RefreshToken string
var LongRefreshToken string
var ClientID string
var ClientSecret string
var CodeAuth string
var RedirectUri string
var Subdomain string
var CrmPhoneField string
var ExolveAccessToken string

func InitSettings() {

	ClientID = FindSettings("ClientID")
	ClientSecret = FindSettings("ClientSecret")
	CodeAuth = FindSettings("CodeAuth")
	RedirectUri = FindSettings("RedirectUri")
	Subdomain = FindSettings("Subdomain")
	CrmPhoneField = FindSettings("CrmPhoneField")
	ExolveAccessToken = FindSettings("ExolveAccessToken")

}

func FindSettings(substr string) string {
	var substring string
	file, err := os.ReadFile("settings/settings")
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
				//fmt.Println("Substring:", substring)
			} else {
				fmt.Println("No match found")
			}
		}
	}
	return substring
}
