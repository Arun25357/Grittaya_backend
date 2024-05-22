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

type NewProduct struct {
	ID          uuid.UUID `jorm:"id"`
	Name        string    `jorm:"product_name"`
	Amount      int       `jorm:"product_amount"`
	UnitPrice   int       `jorm:"product_unitprice"`
	Type        string    `jorm:"product_type"`
	Category    string    `jorm:"product_category"`
	Description string    `jorm:"product_description"`
}
