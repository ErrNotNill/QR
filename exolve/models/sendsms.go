package models

import "time"

type IncomingMessage struct {
	EventID        string    `json:"event_id"`
	MessageID      string    `json:"message_id"`
	ApplicationID  string    `json:"application_id"`
	Date           time.Time `json:"date"`
	Sender         string    `json:"sender"`
	Receiver       string    `json:"receiver"`
	Text           string    `json:"text"`
	Direction      string    `json:"direction"`
	SegmentsCount  int       `json:"segments_count"`
	BillingStatus  string    `json:"billing_status"`
	DeliveryStatus string    `json:"delivery_status"`
	MessageChannel string    `json:"message_channel"`
	Status         string    `json:"status"`
}
