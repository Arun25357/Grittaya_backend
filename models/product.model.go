package models

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(55);not null"`
	Amount      int       `gorm:"not null"`
	UnitPrice   int       `gorm:"not null"`
	Type        string    `gorm:"type:varchar(15);not null"`
	Category    string    `gorm:"type:varchar(15);not null"`
	Description string    `gorm:"type:varchar(15);not null"`
}

type CreateProduct struct {
	ID          uuid.UUID
	Name        string `jorm:"type:varchar(55);not null"`
	Amount      int    `jorm:"not null"`
	UnitPrice   int    `jorm:"not null"`
	Type        string `jorm:"type:varchar(15);not null"`
	Category    string `jorm:"type:varchar(15);not null"`
	Description string `jorm:"type:varchar(15);not null"`
}

type UpdateProduct struct {
	Name        string `jorm:"type:varchar(55);not null"`
	Amount      int    `jorm:"not null"`
	UnitPrice   int    `jorm:"not null"`
	Category    string `jorm:"type:varchar(15);not null"`
	Description string `jorm:"type:varchar(15);not null"`
}

type DeleteProduct struct {
	ID          uuid.UUID
	Name        string `jorm:"type:varchar(55);not null"`
	Amount      int    `jorm:"not null"`
	UnitPrice   int    `jorm:"not null"`
	Type        string `jorm:"type:varchar(15);not null"`
	Category    string `jorm:"type:varchar(15);not null"`
	Description string `jorm:"type:varchar(15);not null"`
}
