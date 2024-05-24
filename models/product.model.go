package models

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(55);not null"`
	Amount      int       `gorm:"not null"`
	UnitPrice   float64   `gorm:"not null"`
	Type        string    `gorm:"type:varchar(55);not null"`
	Category    string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	AttachFile  string
}

type CreateProduct struct {
	Name       string  `json:"product_name" binding:"required"`
	Amount     int     `json:"product_amount" binding:"required"`
	UnitPrice  float64 `json:"product_unitprice" binding:"required"`
	Type       string  `json:"product_type" binding:"required"`
	Category   string  `json:"product_category" binding:"required"`
	AttachFile string  `json:"attach_file"`
}

type UpdateProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	UnitPrice   float64   `json:"product_unitprice"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
	AttachFile  string    `json:"attach_file"`
}

type GetProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	UnitPrice   float64   `json:"product_unitprice"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
	AttachFile  string    `json:"attach_file"`
}

type DeleteProduct struct {
	ID uuid.UUID `json:"product_id"`
}
