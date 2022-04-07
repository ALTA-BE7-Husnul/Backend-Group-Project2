package cart

import (
	"errors"
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

	// ur.database.Exec("UPDATE products SET qty = ? WHERE id = ?", gorm.Expr("qty - ?", cart.Quantity), cart.Product_ID)

	tx := ur.database.Where("product_id = ?", cart.Product_ID).Find(&cartdb)

	if tx.Error != nil {
		return _entities.Cart{}, 0, tx.Error
	}

	if tx.RowsAffected == 1 {
		ur.database.Exec("UPDATE carts SET quantity = ? WHERE product_id = ?", gorm.Expr("quantity + ?", cart.Quantity), cart.Product_ID)
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
	return carts, nil
}
func (ur *CartRepository) PutCart(cart _entities.Cart, idToken int) (_entities.Cart, error) {
	var carts _entities.Cart
	ur.database.Where("product_id = ?", cart.Product_ID).Find(&cart)

	ur.database.Exec("UPDATE carts SET quantity = ? WHERE product_id = ?", gorm.Expr("quantity = ?", cart.Quantity), cart.Product_ID)

	return carts, nil
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
