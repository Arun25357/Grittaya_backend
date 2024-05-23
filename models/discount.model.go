package models

import (
	uuid "github.com/satori/go.uuid"
)

type Discount struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Baht    int       `gorm:"not null"`
	Percent int       `gorm:"not null"`
}
