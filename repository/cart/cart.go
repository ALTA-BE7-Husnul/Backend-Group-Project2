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

	cart.Total = product.Price * uint(cart.Qty)

	if product.Qty < cart.Qty {
		return _entities.Cart{}, 1, errors.New("")
	}

	tx := ur.database.Where("product_id = ?", cart.Product_ID).Find(&cartdb)

	if tx.Error != nil {
		return _entities.Cart{}, 0, tx.Error
	}

	if tx.RowsAffected == 1 {
		ur.database.Exec("UPDATE carts SET qty = ? WHERE product_id = ?", gorm.Expr("qty + ?", cart.Qty), cart.Product_ID)
		ur.database.Exec("UPDATE carts SET total = ? WHERE product_id = ?", gorm.Expr("total + ?", cart.Total), cart.Product_ID)
		return cart, 0, nil
	}

	txSave := ur.database.Create(&cart)
	if txSave.Error != nil {
		return cart, 2, txSave.Error
	}
	return cart, 0, nil

}
func (ur *CartRepository) GetAll() ([]_entities.Cart, error) {
	var carts []_entities.Cart
	tx := ur.database.Find(&carts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(carts)
	return carts, nil
}
func (ur *CartRepository) PutCart(cart _entities.Cart, productId int) (_entities.Cart, error) {
	var product _entities.Product
	txProd := ur.database.Where("id = ?", productId).Find(&product)

	if txProd.Error != nil {
		return _entities.Cart{}, fmt.Errorf("faileed")
	}

	cart.Total = product.Price * cart.Quantity

	tx := ur.database.Where("product_id = ?", productId).Updates(&cart)
	if tx.Error != nil {
		return _entities.Cart{}, fmt.Errorf("faileed")
	}
	return cart, nil
}

func (ur *CartRepository) DeleteCart(id int) (_entities.Cart, int, error) {
	var cart _entities.Cart

	tx := ur.database.Where("product_id = ?", id).Delete(&cart)
	if tx.Error != nil {
		return cart, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, 0, nil
	}
	return cart, int(tx.RowsAffected), nil
}
