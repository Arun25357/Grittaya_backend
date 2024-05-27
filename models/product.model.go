package models

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(55);not null"`
	Amount      int       `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Type        string    `gorm:"type:varchar(55);not null"`
	Category    string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
}

type CreateProduct struct {
	Name       string  `json:"product_name" binding:"required"`
	Amount     int     `json:"product_amount" binding:"required"`
	Price      float64 `json:"product_price" binding:"required"`
	Type       string  `json:"product_type" binding:"required"`
	Category   string  `json:"product_category" binding:"required"`
}

type UpdateProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	Price       float64   `json:"product_Price"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
}

type GetProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	Price       float64   `json:"product_Price"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
}

type DeleteProduct struct {
	ID uuid.UUID `json:"product_id"`
}
