package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Method          string `json:"method" form:"method"`
	DestinationBank string `json:"destination_bank" form:"destinantion_bank"`
}
