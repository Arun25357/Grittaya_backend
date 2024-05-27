package models

import uuid "github.com/satori/go.uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `gorm:"type:varchar(55);not null"`
	Phone    string    `gorm:"type:varchar(10);not null;primary_key"`
	Location string    `gorm:"not null"`
	Postcode int       `gorm:"not null"`
}

type CustomerDetails struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	Postcode int    `json:"postcode"`
}


