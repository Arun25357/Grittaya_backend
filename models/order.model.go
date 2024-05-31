package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Order represents an order in the database
type Order struct {
	ID               uint          `gorm:"type:autoIncrement;primaryKey;uniqueIndex"`
	OrderDate        time.Time     `gorm:"not null"`
	Status           int           `gorm:"not null"`
	CustomerName     string        `gorm:"type:varchar(255);not null"`
	Location         string        `gorm:"not null"`
	DeliveryType     int           `gorm:"not null"`
	TotalPrice       int           `gorm:"not null"`
	Discount         string        `gorm:"type:varchar(55);"`
	SetProductID     uuid.UUID     `gorm:"type:uuid;not null"`
	CustomerID       uuid.UUID     `gorm:"type:uuid;not null"`
	UserID           uuid.UUID     `gorm:"type:uuid;not null"`
	Phone            string        `gorm:"type:varchar(10)"`
	Postcode         int           `gorm:"not null"`
	Platform         string        `gorm:"not null"`
	SetProductName   string        `gorm:"type:varchar(55);not null"`
	Amount           int           `gorm:"not null"`
	Type             string        `gorm:"type:varchar(55);not null"`
	Price            float64       `gorm:"not null"`
	PaymentType      int           `gorm:"not null"`
	LastPricePayment float64       `gorm:"not null"`
	ListProducts     []OrderDetail `gorm:"foreignKey:OrderID"`
}
type CreateOrder struct {
	OrderDate time.Time `json:"order_date"`
	Status    int       `json:"status"`
	UserID    uuid.UUID `json:"user_id"`

	//customer
	CustomerID   uuid.UUID      `json:"customer_id"`
	CustomerName string         `json:"customer_name"`
	Location     string         `json:"location"`
	Postcode     int            `json:"postcode"`
	Phone        string         `json:"phone"`
	Platform     string         `json:"platform"`
	ListProducts []ListProducts `json:"list_products"`

	//Setproduct
	SetProductID   uuid.UUID `json:"set_product_id"`
	SetProductName string    `json:"set_product_name"`
	Amount         int       `json:"set_product_amount"`
	Type           string    `json:"set_product_type"`
	Price          float64   `json:"set_product_price"`
	Discount       string    `json:"discount"` // Adjust if necessary
	TotalPrice     int       `json:"total_price"`

	//Delivery
	DeliveryType int `json:"delivery_type"`

	//Paymemt
	PaymentType      int     `json:"payment_type"`
	LastPricePayment float64 `json:"last_price_payment"`
}
type ListProducts struct {
	SetProductID   uuid.UUID `json:"set_product_id"`
	SetProductName string    `json:"set_product_name"`
	Amount         int       `json:"set_product_amount"`
	Type           string    `json:"set_product_type"`
	Price          float64   `json:"set_product_price"`
	Discount       string    `json:"discount"`
	TotalPrice     int       `json:"total_price"`
}

type UpdateOrder struct {
	ID        uint      `json:"order_id" uri:"id"`
	OrderDate time.Time `json:"order_date"`
	Status    int       `json:"status"`
	UserID    uuid.UUID `json:"user_id"`

	//Update customer
	CustomerID   uuid.UUID `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Location     string    `json:"location"`
	Postcode     int       `json:"postcode"`
	Phone        string    `json:"phone"`
	Platform     string    `json:"platform"`

	//Setproduct
	SetProductID   uuid.UUID      `json:"set_product_id"`
	SetProductName string         `json:"set_product_name"`
	Amount         int            `json:"set_product_amount"`
	Type           string         `json:"set_product_type"`
	Price          float64        `json:"set_product_price"`
	Discount       string         `json:"discount"`
	TotalPrice     int            `json:"total_price"`
	ListProducts   []ListProducts `json:"list_products"`

	//Delivery
	DeliveryType int `json:"delivery_type"`

	//Paymemt
	PaymentType      int     `json:"payment_type"`
	LastPricePayment float64 `json:"last_price_payment"`
}

// GetOrder represents the payload for retrieving an order
type GetOrder struct {
	ID        uint      `json:"order_id"`
	OrderDate time.Time `json:"order_date"`
	Status    int       `json:"status"`
	UserID    uuid.UUID `json:"user_id"`

	//Getorder customer
	CustomerID   uuid.UUID `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Location     string    `json:"location"`
	Postcode     int       `json:"postcode"`
	Phone        string    `json:"phone"`
	Platform     string    `json:"platform"`

	//Getorder Setproduct
	SetProductID   uuid.UUID      `json:"set_product_id"`
	SetProductName string         `json:"set_product_name"`
	Amount         int            `json:"set_product_amount"`
	Type           string         `json:"set_product_type"`
	Price          float64        `json:"set_product_price"`
	Discount       string         `json:"discount"`
	TotalPrice     int            `json:"total_price"`
	ListProducts   []ListProducts `json:"list_products"`

	//Getorder Delivery
	DeliveryType int `json:"delivery_type"`

	//Getorder Paymemt
	PaymentType      int     `json:"payment_type"`
	LastPricePayment float64 `json:"last_price_payment"`
}

// DeleteOrder represents the payload for deleting an order
type DeleteOrder struct {
	ID uint `json:"id" binding:"required"`
}

type RequestOrderCreate struct {
	OrderID          int       `gorm:"not null"`
	SetProductID     uuid.UUID `gorm:"type:uuid;not null"`
	Amount           int       `gorm:"not null"`
}
// type RequestOrderUpdate struct {

// }
// type RequestOrderGetByID struct {

// }
// type RequestOrderGetList struct {

// }
// type ResponseOrderGetList struct {

// }
