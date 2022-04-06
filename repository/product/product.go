package product

import (
	"errors"
	_entities "group-project-2/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: db,
	}
}

// add product
func (pr *ProductRepository) AddProduct(product _entities.Product) (_entities.Product, error) {
	tx := pr.database.Save(&product)
	if tx.Error != nil {
		return product, tx.Error
	}
	return product, nil
}

// update product
func (pr *ProductRepository) UpdateProductById(product _entities.Product, id, idToken int) (_entities.Product, int, error) {
	if product.Name != "" {
		tx := pr.database.Model(&_entities.Product{}).Where("id = ?", id).Where("seller_id = ?", idToken).Update("name", product.Name)
		if tx.Error != nil {
			return product, 0, tx.Error
		}
		if tx.RowsAffected == 0 {
			return product, 0, tx.Error
		}
	}
	if product.Category != "" {
		tx := pr.database.Model(&_entities.Product{}).Where("id = ?", id).Where("seller_id = ?", idToken).Update("category", product.Category)
		if tx.Error != nil {
			return product, 0, tx.Error
		}
		if tx.RowsAffected == 0 {
			return product, 0, tx.Error
		}
	}
	if product.Price != 0 {
		tx := pr.database.Model(&_entities.Product{}).Where("id = ?", id).Where("seller_id = ?", idToken).Update("price", product.Price)
		if tx.Error != nil {
			return product, 0, tx.Error
		}
		if tx.RowsAffected == 0 {
			return product, 0, tx.Error
		}
	}
	if product.Qty != 0 {
		tx := pr.database.Model(&_entities.Product{}).Where("id = ?", id).Where("seller_id = ?", idToken).Update("qty", product.Qty)
		if tx.Error != nil {
			return product, 0, tx.Error
		}
		if tx.RowsAffected == 0 {
			return product, 0, tx.Error
		}
	}
	if product.Description != "" {
		tx := pr.database.Model(&_entities.Product{}).Where("id = ?", id).Where("seller_id = ?", idToken).Update("description", product.Description)
		if tx.Error != nil {
			return product, 0, tx.Error
		}
		if tx.RowsAffected == 0 {
			return product, 0, tx.Error
		}
	}
	return product, 1, nil
}

// delete by id
func (pr *ProductRepository) DeleteProductById(id, idToken int) (int, error) {
	var product _entities.Product
	tx := pr.database.Where("id = ?", id).Where("seller_id = ?", idToken).Delete(&product)
	if tx.Error != nil {
		return 0, tx.Error
	}
	x := tx.RowsAffected
	if x == 0 {
		return 0, errors.New("error")
	}
	return int(x), nil
}

// get all
func (pr *ProductRepository) GetAllProduct() ([]_entities.Product, error) {
	var products []_entities.Product
	tx := pr.database.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}

// get by id
func (pr *ProductRepository) GetProductById(id int) (_entities.Product, int, error) {
	var product _entities.Product
	tx := pr.database.Where("id = ?", id).Find(&product)
	if tx.Error != nil {
		return product, 0, tx.Error
	}
	x := tx.RowsAffected
	if x == 0 {
		return product, 0, errors.New("product not found")
	}
	return product, int(x), nil
}
