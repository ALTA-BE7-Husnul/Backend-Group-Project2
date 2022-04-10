package helper

type OrderRequestFormat struct {
	Cart_ID []uint               `json:"cart_id" form:"cart_id"`
	Address AddressRequestFormat `json:"address" form:"address"`
	Payment PaymentRequestFormat `json:"payment" form:"payment"`
}

type AddressRequestFormat struct {
	Street  string `json:"street" form:"street"`
	City    string `json:"city" form:"city"`
	State   string `json:"state" form:"state"`
	Zipcode uint   `json:"zipcode" form:"zipcode"`
}

type PaymentRequestFormat struct {
	Method          string `json:"method" form:"method"`
	DestinationBank string `json:"destination_bank" form:"destinantion_bank"`
}
