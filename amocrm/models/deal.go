package models

type CustomFieldValue struct {
	FieldID int `json:"field_id"`
	Values  []struct {
		Value string `json:"value"`
	} `json:"values"`
}

type Embedded struct {
	Tags []struct {
		ID int `json:"id"`
	} `json:"tags"`
}

type Lead struct {
	Name              string             `json:"name"`
	CreatedBy         int                `json:"created_by,omitempty"`
	Price             int                `json:"price"`
	CustomFieldValues []CustomFieldValue `json:"custom_fields_values,omitempty"`
	Embedded          Embedded           `json:"_embedded,omitempty"`
}
