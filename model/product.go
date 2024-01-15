package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	ID			uuid.UUID	`gorm:"type:uuid;"`
	Name		string		`json:"name"`
	Price		int			`json:"price"`
	URLImage	string		`json:"url_image"`
}

// Products struct
type Products struct {
 	Products []Product `json:"products"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	product.ID = uuid.New()
	return
}