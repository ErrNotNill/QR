package models

import "time"

type Message struct {
	MessageID       string    `json:"message_id"`
	ApplicationUUID string    `json:"application_uuid"`
	Date            time.Time `json:"date"`
	Number          string    `json:"number"`
	Sender          string    `json:"sender"`
	Receiver        string    `json:"receiver"`
	Text            string    `json:"text"`
	Direction       int       `json:"direction"`
	SegmentsCount   int       `json:"segments_count"`
	BillingStatus   int       `json:"billing_status"`
	DeliveryStatus  int       `json:"delivery_status"`
	Channel         int       `json:"channel"`
	Status          int       `json:"status"`
}

type MessageResponse struct {
	Messages []Message `json:"messages"`
}

type Count struct {
	Count string `json:"count"`
}
