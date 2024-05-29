package models

import (
	uuid "github.com/satori/go.uuid"
)

type SetProduct struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(55);not null"`
	Amount    int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Type      string    `gorm:"type:varchar(55);not null"`
	Status    int       `gorm:"not null"`
}

type CreateSetProduct struct {
	Name   string  `json:"set_product_name" binding:"required"`
	Amount int     `json:"set_product_amount" binding:"required"`
	Price  float64 `json:"set_product_price" binding:"required"`
	Status int     `json:"status" binding:"required"`
	Type   string  `json:"set_product_type"`
}

type UpdateSetProduct struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Name      string    `json:"set_product_name"`
	Amount    int       `json:"set_product_amount"`
	Price     float64   `json:"set_product_price"`
	Status    int       `json:"status"`
	Type      string    `json:"set_product_type"`
}

type GetSetProduct struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"set_product_name"`
	Amount    int       `json:"set_product_amount"`
	Price     float64   `json:"set_product_price"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int       `json:"status"`
	Type      string    `json:"set_product_type"`
}

type DeleteSetProduct struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
