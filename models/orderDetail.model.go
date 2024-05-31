package models

import (
	uuid "github.com/satori/go.uuid"
)

type OrderDetail struct {
	ID           uint      `gorm:"type:autoIncrement;primaryKey;uniqueIndex"`
	OrderID      int       `gorm:"not null"`
	SetProductID uuid.UUID `gorm:"type:uuid;not null"`
	Amount       int       `gorm:"not null"`
}
type DeleteOrderDetail struct {
	ID uint `json:"orderdetail_id"`
}
