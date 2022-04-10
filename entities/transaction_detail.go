package entities

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	// User_ID        uint        `json:"user_id" form:"user_id"`
	Product_ID     uint    `json:"product_id" form:"product_id"`
	Transaction_ID uint    `json:"transaction_id" form:"transaction_id"`
	Status         string  `gorm:"default:not paid" json:"status" form:"status"`
	Total          uint    `json:"total" form:"total"`
	Product        Product `gorm:"foreignKey:Product_ID;references:ID"`
}
