package models

import "github.com/satori/go.uuid"

type Detailorder struct {
	ID       		uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	DeliveryType	string 		`gorm:"type:varchar(55);uniqueIndex;not null"`
	TotalPrice		int 		`gorm:"not null"`
	PaymentType		string		`gorm:"type:varchar(55);uniqueIndex;not null"`
	Discount		string		`gorm:"type:varchar(55);uniqueIndex;not null"`
	OrderID			uuid.UUID 	`gorm:"type:varchar(255);not null"`
	SetproductID	uuid.UUID 	`gorm:"type:varchar(255);not null"`
	CustomerID		uuid.UUID 	`gorm:"type:varchar(255);not null"`
}
