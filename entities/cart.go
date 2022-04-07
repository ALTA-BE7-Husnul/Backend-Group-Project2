package entities

type Cart struct {
	Buyer_ID   uint   `json:"buyer_id" from:"buyer_id"`
	Product_ID uint   `json:"product_id" from:"product_id"`
	Quantit    uint   `json:"quantity" from:"quantity"`
	Total      uint   `json:"total" from:"total"`
	Status     string `json:"status" from:"status"`
}
