package models

import uuid "github.com/satori/go.uuid"

type Setproduct struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(55);not null"`
	Amount    int       `gorm:"not null"`
	UnitPrice int       `gorm:"not null"`
	ProductID uuid.UUID `gorm:"type:varchar(255);not null"`
	Status    int       `gorm:"not null"`
}

type SetProduct struct {
	ID          uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string           `gorm:"type:varchar(255);not null"`
	Description string           `gorm:"type:varchar(255);not null"`
	Items       []SetProductItem `gorm:"foreignKey:SetProductID"`
}

// Model to associate products with sets
type SetProductItem struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	SetProductID uuid.UUID `gorm:"type:uuid;not null"`
	ProductID    uuid.UUID `gorm:"type:uuid;not null"`
	Product      Product   `gorm:"foreignKey:ProductID"`
	Quantity     int       `gorm:"not null"`
}
