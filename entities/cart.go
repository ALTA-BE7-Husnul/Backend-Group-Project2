package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Buyer_ID   uint `gorm:"not null" json:"buyer_id" from:"buyer_id"`
	Product_ID uint `gorm:"not null" json:"product_id" from:"product_id"`
	Qty        int  `gorm:"not null" json:"qty" from:"qty"`
	Total      uint `gorm:"not null" json:"total" from:"total"`
	// Product    uint `gorm:"foreignKey:Product_ID;references:ID" json:"product" form:"product"`
}
