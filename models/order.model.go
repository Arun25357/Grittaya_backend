package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDate        time.Time
	Status           int
	CustomerUsername string    `gorm:"type:varchar(255);not null"`
	Platform         string    `gorm:"type:varchar(255);not null"`
	DeliveryType     string    `gorm:"type:varchar(255);not null"`
	TotalPrice       int       `gorm:"type:int;not null"`
	UserID          uuid.UUID `gorm:"type:varchar(255);not null"`
}
