package models

import uuid "github.com/satori/go.uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `gorm:"type:varchar(55);not null"`
	Phone    string    `gorm:"type:varchar(12);not null;primary_key"`
	Location string    `gorm:"not null"`
	Postcode int       `gorm:"not null"`
	Platform string    `gorm:"not null"`
}

type CustomerDetails struct {
	Name     string `json:"customer_name"`
	Phone    string `json:"customer_phone"`
	Location string `json:"customer_location"`
	Postcode int    `json:"customer_postcode"`
	Platform string `json:"customer_platform"`
}
