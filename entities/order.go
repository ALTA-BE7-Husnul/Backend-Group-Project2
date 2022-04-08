package entities

import (
	_helper "group-project-2/delivery/helper"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	User_ID uint   `json:"user_id" form:"user_id"`
	Cart_ID []uint `json:"cart_id" form:"cart_id"`
	// Total   uint   `json:"total" form:"total"`
	// Status  string                       `json:"status" form:"status"`
	Address _helper.AddressRequestFormat `json:"address" form:"address"`
	Payment _helper.PaymentRequestFormat `json:"payment" form:"payment"`
	// TransactionDetail []Transaction `gorm:"foreignkey:Transaction_ID"`
	// Address []Address `gorm:"foreignkey:Transaction_ID"`
	// Payment []Payment `gorm:"foreignkey:Transaction_ID"`
}

// type TransactionDetail struct {
// 	gorm.Model
// 	Transaction_ID uint   `json:"transaction_id" form:"transaction_id"`
// 	Cart_ID        uint   `json:"cart_id" form:"cart_id"`
// 	Status         string `json:"status" form:"status"`
// }

// type Address struct {
// 	gorm.Model
// 	Transaction_ID uint   `gorm:"primarykey" json:"transaction_id" form:"transaction_id"`
// 	Street         string `json:"street" form:"street"`
// 	City           string `json:"city" form:"city"`
// 	State          string `json:"state" form:"state"`
// 	Zipcode        uint   `json:"zipcode" form:"zipcode"`
// }

// type Payment struct {
// 	gorm.Model
// 	Transaction_ID  uint   `gorm:"primarykey" json:"transaction_id" form:"transaction_id"`
// 	Method          string `json:"method" form:"method"`
// 	DestinationBank string `json:"destination_bank" form:"destinantion_bank"`
// }
