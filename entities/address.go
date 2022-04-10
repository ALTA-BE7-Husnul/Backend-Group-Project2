package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street  string `json:"street" form:"street"`
	City    string `json:"city" form:"city"`
	State   string `json:"state" form:"state"`
	Zipcode uint   `json:"zipcode" form:"zipcode"`
}
