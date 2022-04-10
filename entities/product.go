package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Seller_ID   uint   `json:"seller_id" form:"seller_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
	Price       uint   `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Image       string `json:"image" form:"image"`
	Cart        []Cart `gorm:"foreignkey:Product_ID;references:ID"`
}
