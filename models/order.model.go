package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Order represents an order in the database
type Order struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderDate         time.Time `gorm:"not null"`
	Status            int       `gorm:"not null"`
	CustomerUsername  string    `gorm:"type:varchar(255);not null"`
	DeliveryType      int       `gorm:"not null"`
	TotalPrice        int       `gorm:"not null"`
	Discount          string    `gorm:"type:varchar(55);uniqueIndex;not null"` // Adjust as needed
	SetproductID      uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID        uuid.UUID `gorm:"type:uuid;not null"`
	UserID            uuid.UUID `gorm:"type:uuid;not null"`
	Postcode          int       `gorm:"not null"`
	Platform          int       `gorm:"not null"`
	SetproductName    string    `gorm:"type:varchar(55);not null"`
	Amount            int       `gorm:"not null"`
	Type              string    `gorm:"type:varchar(55);not null"`
	Price             float64   `gorm:"not null"`
	PaymentType       int       `gorm:"not null"`
	TotalpricePayment float64   `gorm:"not null"`
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
	Platform         int       `json:"platform"`

	//Setproduct
	SetproductID uuid.UUID `json:"setproduct_id" binding:"required"`
	Name         string    `json:"setproduct_name"`
	Amount       int       `json:"setproduct_amount"`
	Type         string    `json:"setproduct_type"`
	Price        float64   `json:"setproduct_price" binding:"required"`
	Discount     string    `json:"discount"` // Adjust if necessary
	TotalPrice   int       `json:"total_price" binding:"required"`

	//Delivery
	DeliveryType int `json:"delivery_type" binding:"required"`

	//Paymemt
	PaymentType       int     `json:"payment_type" binding:"required"`
	TotalpricePayment float64 `json:"totalprice_payment" binding:"required"`
}

// GetOrder represents the payload for retrieving an order
type GetOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

// DeleteOrder represents the payload for deleting an order
type DeleteOrder struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
