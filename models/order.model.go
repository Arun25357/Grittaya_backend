package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Order struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDate        time.Time
	DeliveryType     string    `gorm:"type:varchar(255);not null"`
	ProductName      string    `gorm:"type:varchar(255);not null"`
	ProductAmount    int       `gorm:"type:int;not null"`
	TotalPrice       int       `gorm:"type:int;not null"`
	AdminID          uuid.UUID `gorm:"type:varchar(255);not null"`
	CustomerUsername string    `gorm:"type:varchar(255);not null"`
}
