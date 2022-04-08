package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Buyer_ID    uint          `json:"buyer_id" from:"buyer_id"`
	Product_ID  uint          `json:"product_id" from:"product_id"`
	Quantity    uint          `json:"quantity" from:"quantity"`
	Total       uint          `json:"total" from:"total"`
	Status      string        `gorm:"default:unpaid" json:"status" from:"status"`
	Transaction []Transaction `gorm:"foreignkey:Cart_ID"`
}
