package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Order represents an order in the database
type Order struct {
	ID               uint      `gorm:"type:autoIncrement;primaryKey;uniqueIndex"`
	OrderDate        time.Time `gorm:"not null"`
	Status           int       `gorm:"not null"`
	CustomerUsername string    `gorm:"type:varchar(255);not null"`
	DeliveryType     int       `gorm:"not null"`
	TotalPrice       int       `gorm:"not null"`
	Discount         string    `gorm:"type:varchar(55);uniqueIndex;not null"` // Adjust as needed
	SetProductID     uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID       uuid.UUID `gorm:"type:uuid;not null"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
	Postcode         int       `gorm:"not null"`
	Platform         string    `gorm:"not null"`
	SetProductName   string    `gorm:"type:varchar(55);not null"`
	Amount           int       `gorm:"not null"`
	Type             string    `gorm:"type:varchar(55);not null"`
	Price            float64   `gorm:"not null"`
	PaymentType      int       `gorm:"not null"`
	LastPricePayment float64   `gorm:"not null"`
}
type CreateOrder struct {
	OrderDate time.Time `json:"order_date" binding:"required"`
	Status    int       `json:"status" binding:"required"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`

	//customer
	CustomerID       uuid.UUID `json:"customer_id" binding:"required"`
	CustomerUsername string    `json:"customer_username" binding:"required"`
	Location         string    `json:"location"`
	Postcode         int       `json:"postcode"`
	Platform         string    `json:"platform"`

	//Setproduct
	SetProductID   uuid.UUID `json:"set_product_id" binding:"required"`
	SetProductName string    `json:"set_product_name"`
	Amount         int       `json:"set_product_amount"`
	Type           string    `json:"set_product_type"`
	Price          float64   `json:"set_product_price" binding:"required"`
	Discount       string    `json:"discount"` // Adjust if necessary
	TotalPrice     int       `json:"total_price" binding:"required"`

	//Delivery
	DeliveryType int `json:"delivery_type" binding:"required"`

	//Paymemt
	PaymentType      int     `json:"payment_type" binding:"required"`
	LastPricePayment float64 `json:"last_price_payment" binding:"required"`
}

// GetOrder represents the payload for retrieving an order
type GetOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

// DeleteOrder represents the payload for deleting an order
type DeleteOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
