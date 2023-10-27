package exolve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"qr/exolve/models"
	"time"
)

func CreateBody(datestart, datefinal string) []byte {
	dateGte, err := time.Parse(time.RFC3339, datestart)
	if err != nil {
		fmt.Println(err)
	}
	dateLte, err := time.Parse(time.RFC3339, datefinal)
	if err != nil {
		fmt.Println(err)
	}
	command := fmt.Sprintf(`{
      "date_gte": "%s",
      "date_lte": "%s"
    }`, dateGte.Format(time.RFC3339), dateLte.Format(time.RFC3339))
	//fmt.Println("command time: ", command)
	return []byte(command)
}

func GetList() {
	uri := `https://api.exolve.ru/messaging/v1/GetList`

	dateStart := "2023-10-20T15:04:05.000Z"  //YYYY-MM-DD
	dateFinish := "2023-10-27T15:04:05.000Z" //YYYY-MM-DD
	body := CreateBody(dateStart, dateFinish)

	AccessToken := `eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRV05sMENiTXY1SHZSV29CVUpkWjVNQURXSFVDS0NWODRlNGMzbEQtVHA0In0.eyJleHAiOjIwMDk3MjkwNDUsImlhdCI6MTY5NDM2OTA0NSwianRpIjoiYTA2NDg5YjItMjc4YS00MWQwLTg5NzktYzU3ZjNmM2NkZWI2IiwiaXNzIjoiaHR0cHM6Ly9zc28uZXhvbHZlLnJ1L3JlYWxtcy9FeG9sdmUiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiYzM5OWY3MmMtOTBmOS00ZTYxLTg1ZjYtMGYxOGUxYjkzYWZkIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiMWJkOGJmYzktNmExMi00Yjg0LTljZmUtYmEzNjNiYmQ1MjNlIiwic2Vzc2lvbl9zdGF0ZSI6ImI1MjliYjVhLTJiNzAtNDIzMS1hYTdiLWVhNzg4NzBlYjA0YSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1leG9sdmUiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJleG9sdmVfYXBwIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJiNTI5YmI1YS0yYjcwLTQyMzEtYWE3Yi1lYTc4ODcwZWIwNGEiLCJ1c2VyX3V1aWQiOiJjOGNjNDgwNC03YmQ5LTQ2NWQtYTVmNC0xMjY1NjA0NGMzMmEiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxNzIuMjAuMi4yMiIsImNsaWVudElkIjoiMWJkOGJmYzktNmExMi00Yjg0LTljZmUtYmEzNjNiYmQ1MjNlIiwiYXBpX2tleSI6dHJ1ZSwiYXBpZm9uaWNhX3NpZCI6IjFiZDhiZmM5LTZhMTItNGI4NC05Y2ZlLWJhMzYzYmJkNTIzZSIsImJpbGxpbmdfbnVtYmVyIjoiMTIwNDA3MyIsImFwaWZvbmljYV90b2tlbiI6ImF1dDRkYjc5MDFjLWMxMWYtNDFlYi1iYWE2LTVkNGI2MDA1MDY0YiIsInByZWZlcnJlZF91c2VybmFtZSI6InNlcnZpY2UtYWNjb3VudC0xYmQ4YmZjOS02YTEyLTRiODQtOWNmZS1iYTM2M2JiZDUyM2UiLCJjdXN0b21lcl9pZCI6IjI4ODQ5IiwiY2xpZW50QWRkcmVzcyI6IjE3Mi4yMC4yLjIyIn0.By1UabvoodzGQVwGFuRy3gv1iuf_YTeY4mbj-bP0LA5FhJ-Bp0TVFGUyv_WnOWNVN59SJbkSh3dxV0Ydo62uq7tKaQd3UF5fmfAHCacuGjH9CNARPgU8UeuR5XwHVDtlOYt9F2wGCSG09NO5x-YGaPJdu_qIAn45g_OV8bLdJhof7jXrS6DgDWWzxroo7D7g2UdpPk0xwU9Brj5Y3kUyxfIp9ZwPTTPnr-MGgJJHgQv8mWvsddKYvxW26MSeR4E3pFsU2_8GtaQXbX_g8rXPpb0yjwTfn_qExZlfG6eQXrOlKAFRUc7wMtZhGdJLbaOW7tBrXptCYmUmULaQApq8bw`

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		log.Println("")
	}
	req.Header.Add("Authorization", "Bearer "+AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var message models.Message
	bs, _ := io.ReadAll(req.Body)
	jsonData := json.Unmarshal(bs, &message)
	if jsonData != nil {
		fmt.Println("Message received successfully")
	}
	log.Println("response :", message)

	log.Println("number phone who send: ", message.Number)
	log.Println("messageID : ", message.MessageID)
}

func GetCount() {
	uri := `https://api.exolve.ru/messaging/v1/GetCount`

	dateStart := "2023-10-20T15:04:05.000Z"  //YYYY-MM-DD
	dateFinish := "2023-10-27T15:04:05.000Z" //YYYY-MM-DD
	body := CreateBody(dateStart, dateFinish)

	AccessToken := `eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRV05sMENiTXY1SHZSV29CVUpkWjVNQURXSFVDS0NWODRlNGMzbEQtVHA0In0.eyJleHAiOjIwMDk3MjkwNDUsImlhdCI6MTY5NDM2OTA0NSwianRpIjoiYTA2NDg5YjItMjc4YS00MWQwLTg5NzktYzU3ZjNmM2NkZWI2IiwiaXNzIjoiaHR0cHM6Ly9zc28uZXhvbHZlLnJ1L3JlYWxtcy9FeG9sdmUiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiYzM5OWY3MmMtOTBmOS00ZTYxLTg1ZjYtMGYxOGUxYjkzYWZkIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiMWJkOGJmYzktNmExMi00Yjg0LTljZmUtYmEzNjNiYmQ1MjNlIiwic2Vzc2lvbl9zdGF0ZSI6ImI1MjliYjVhLTJiNzAtNDIzMS1hYTdiLWVhNzg4NzBlYjA0YSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1leG9sdmUiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJleG9sdmVfYXBwIHByb2ZpbGUgZW1haWwiLCJzaWQiOiJiNTI5YmI1YS0yYjcwLTQyMzEtYWE3Yi1lYTc4ODcwZWIwNGEiLCJ1c2VyX3V1aWQiOiJjOGNjNDgwNC03YmQ5LTQ2NWQtYTVmNC0xMjY1NjA0NGMzMmEiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxNzIuMjAuMi4yMiIsImNsaWVudElkIjoiMWJkOGJmYzktNmExMi00Yjg0LTljZmUtYmEzNjNiYmQ1MjNlIiwiYXBpX2tleSI6dHJ1ZSwiYXBpZm9uaWNhX3NpZCI6IjFiZDhiZmM5LTZhMTItNGI4NC05Y2ZlLWJhMzYzYmJkNTIzZSIsImJpbGxpbmdfbnVtYmVyIjoiMTIwNDA3MyIsImFwaWZvbmljYV90b2tlbiI6ImF1dDRkYjc5MDFjLWMxMWYtNDFlYi1iYWE2LTVkNGI2MDA1MDY0YiIsInByZWZlcnJlZF91c2VybmFtZSI6InNlcnZpY2UtYWNjb3VudC0xYmQ4YmZjOS02YTEyLTRiODQtOWNmZS1iYTM2M2JiZDUyM2UiLCJjdXN0b21lcl9pZCI6IjI4ODQ5IiwiY2xpZW50QWRkcmVzcyI6IjE3Mi4yMC4yLjIyIn0.By1UabvoodzGQVwGFuRy3gv1iuf_YTeY4mbj-bP0LA5FhJ-Bp0TVFGUyv_WnOWNVN59SJbkSh3dxV0Ydo62uq7tKaQd3UF5fmfAHCacuGjH9CNARPgU8UeuR5XwHVDtlOYt9F2wGCSG09NO5x-YGaPJdu_qIAn45g_OV8bLdJhof7jXrS6DgDWWzxroo7D7g2UdpPk0xwU9Brj5Y3kUyxfIp9ZwPTTPnr-MGgJJHgQv8mWvsddKYvxW26MSeR4E3pFsU2_8GtaQXbX_g8rXPpb0yjwTfn_qExZlfG6eQXrOlKAFRUc7wMtZhGdJLbaOW7tBrXptCYmUmULaQApq8bw`

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		log.Println("")
	}
	req.Header.Add("Authorization", "Bearer "+AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))
}
