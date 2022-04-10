package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string        `json:"name" form:"name"`
	Email       string        `gorm:"unique" json:"email" form:"email"`
	Address     string        `json:"address" form:"address"`
	Password    string        `json:"password" form:"password"`
	Product     []Product     `gorm:"foreignkey:Seller_ID;references:ID"`
	Cart        []Cart        `gorm:"foreignkey:Buyer_ID;references:ID"`
	Transaction []Transaction `gorm:"foreignkey:User_ID;references:ID"`
}
