package models

import (
	"time"

	"github.com/satori/go.uuid"
)

// Order represents an order in the database
type Order struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDate        time.Time `gorm:"not null"`
	Status           int       `gorm:"not null"`
	CustomerUsername string    `gorm:"type:varchar(255);not null"`
	Platform         string    `gorm:"type:varchar(255);not null"`
	DeliveryType     string    `gorm:"type:varchar(255);not null"`
	TotalPrice       int       `gorm:"not null"`
	Discount         string    `gorm:"type:varchar(55);uniqueIndex;not null"` // Adjust as needed
	SetproductID     uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID       uuid.UUID `gorm:"type:uuid;not null"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
}

// CreateOrder represents the payload for creating an order
type CreateOrder struct {
	OrderDate        time.Time `json:"order_date" binding:"required"`
	Status           int       `json:"status" binding:"required"`
	CustomerUsername string    `json:"customer_username" binding:"required"`
	Platform         string    `json:"platform" binding:"required"`
	DeliveryType     string    `json:"delivery_type" binding:"required"`
	TotalPrice       int       `json:"total_price" binding:"required"`
	Discount         string    `json:"discount"` // Adjust if necessary
	SetproductID     uuid.UUID `json:"setproduct_id" binding:"required"`
	CustomerID       uuid.UUID `json:"customer_id" binding:"required"`
	UserID           uuid.UUID `json:"user_id" binding:"required"`
}

// GetOrder represents the payload for retrieving an order
type GetOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

// DeleteOrder represents the payload for deleting an order
type DeleteOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
