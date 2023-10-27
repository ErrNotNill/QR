package models

type Direction int

const (
	DIRECTION_INCOMING Direction = 1
	DIRECTION_OUTGOING Direction = 2
)

type BillingStatus int

const (
	BILLING_STATUS_PREBILLED   BillingStatus = 1
	BILLING_STATUS_BILLED      BillingStatus = 2
	BILLING_STATUS_UNDERFUNDED BillingStatus = 4
	BILLING_STATUS_FAILED      BillingStatus = 6
	BILLING_STATUS_AUTHORIZED  BillingStatus = 7
)

type DeliveryStatus int

const (
	DELIVERY_STATUS_QUEUED           DeliveryStatus = 1
	DELIVERY_STATUS_TRANSMITTED      DeliveryStatus = 2
	DELIVERY_STATUS_DELIVERED        DeliveryStatus = 3
	DELIVERY_STATUS_FAILED           DeliveryStatus = 4
	DELIVERY_STATUS_RETRIES_EXCEEDED DeliveryStatus = 5
	DELIVERY_STATUS_PROHIBITED       DeliveryStatus = 6
)

type Status int

const (
	STATUS_QUEUED      Status = 1
	STATUS_TRANSMITTED Status = 2
	STATUS_DELIVERED   Status = 3
	STATUS_FAILED      Status = 4
	STATUS_UNDERFUNDED Status = 5
	STATUS_PROHIBITED  Status = 6
)

type Message struct {
	Messages        []string
	MessageID       int64
	ApplicationUUID string
	Date            string
	Number          string
	Sender          string
	Receiver        string
	Text            string
	Direction       Direction
	SegmentsCount   uint32
	BillingStatus   BillingStatus
	DeliveryStatus  DeliveryStatus
	Channel         string // Enum type for the channel
	Status          Status // Enum type for the status
}
