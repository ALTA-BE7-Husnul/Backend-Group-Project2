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

func (ur *CartRepository) PostCart(cart _entities.Cart, idToken int) (_entities.Cart, int, error) {
	var product _entities.Product
	var cartdb _entities.Cart
	cart.Buyer_ID = uint(idToken)

	ur.database.Where("id = ?", cart.Product_ID).Find(&product)

	cart.Total = product.Price * cart.Quantity

	if product.Qty < cart.Quantity {
		return _entities.Cart{}, 1, errors.New("")
	}

	product.Qty = product.Qty - cart.Quantity

	ur.database.Exec("UPDATE products SET qty = ? WHERE id = ?", gorm.Expr("qty - ?", cart.Quantity), cart.Product_ID)

	tx := ur.database.Where("product_id = ?", cart.Product_ID).Find(&cartdb)

	if tx.Error != nil {
		return _entities.Cart{}, 0, tx.Error
	}
	fmt.Println("ini row aff", tx.RowsAffected)
	if tx.RowsAffected == 1 {
		fmt.Println("ini quantiti ", cart.Quantity)
		ur.database.Exec("UPDATE carts SET quantity = ? WHERE product_id = ?", gorm.Expr("quantity + ?", cart.Quantity), cart.Product_ID)
		ur.database.Exec("UPDATE carts SET total = ? WHERE product_id = ?", gorm.Expr("total + ?", cart.Total), cart.Product_ID)
		return cart, 0, nil
	}
	fmt.Println("lewat")
	txSave := ur.database.Create(&cart)
	if txSave.Error != nil {
		return cart, 2, txSave.Error
	}
	return cart, 0, nil

}
