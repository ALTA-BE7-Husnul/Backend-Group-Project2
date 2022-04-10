package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	User_ID           uint                `json:"user_id" form:"user_id"`
	Total             uint                `json:"total" form:"total"`
	Status            string              `gorm:"default:not paid" json:"status" form:"status"`
	TransactionDetail []TransactionDetail `gorm:"foreignkey:Transaction_ID;references:ID"`
	Address           Address             `gorm:"foreignkey:ID;references:ID" json:"address" form:"payment"`
	Payment           Payment             `gorm:"foreignkey:ID;references:ID" json:"payment" form:"payment"`
}
