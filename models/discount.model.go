package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Discount struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Percent int       `gorm:"not null"`
	Expir   time.Time
}
