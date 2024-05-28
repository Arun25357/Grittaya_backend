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
	Status    int       `gorm:"not null"`
}

type CreateSetProduct struct {
	Name      string    `json:"name" binding:"required"`
	Amount    int       `json:"amount" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Status    int       `json:"status" binding:"required"`
}

type UpdateSetProduct struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	Price     float64   `json:"price"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int       `json:"status"`
}

type GetSetProduct struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	Price     float64   `json:"price"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int       `json:"status"`
}

type DeleteSetProduct struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
