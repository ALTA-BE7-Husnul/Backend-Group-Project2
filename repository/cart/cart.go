package cart

import (
	"errors"
	"fmt"
	_entities "group-project-2/entities"

	"gorm.io/gorm"
)

type CartRepository struct {
	database *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		database: db,
	}
}

func (ur *CartRepository) PostCart(cart _entities.Cart, idToken int) (_entities.Cart, error, int) {
	var product _entities.Product
	cart.Buyer_ID = uint(idToken)
	ur.database.Where("id = ?", cart.Product_ID).Find(&product)
	fmt.Println("ini data yang di dapat", product.Name)
	cart.Total = product.Price * cart.Quantity

	product.Qty = product.Qty - cart.Quantity

	if product.Qty < cart.Quantity {
		return _entities.Cart{}, errors.New("Not enough product"), 1
	}

	ur.database.Exec("UPDATE products SET qty = ? WHERE id = ?", gorm.Expr("qty - ?", cart.Quantity), cart.Product_ID)
	// ur.database.Where("ID = ?", cart.Product_ID).Updates(products)
	tx := ur.database.Save(&cart)
	if tx.Error != nil {
		return cart, tx.Error, 2
	}
	return cart, nil, 0

}
