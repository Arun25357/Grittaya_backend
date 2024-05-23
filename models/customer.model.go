package models

import uuid "github.com/satori/go.uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);not null"`
	Phone    string    `gorm:"type:varchar(15);not null"`
	Location string    `gorm:"not null"`
	Postcode int       `gorm:"not null"`
}
