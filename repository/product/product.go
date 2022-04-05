package product

import (
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

func (pr *ProductRepository) AddProduct(product _entities.Product) (_entities.Product, error) {
	tx := pr.database.Save(&product)
	if tx.Error != nil {
		return product, tx.Error
	}
	return product, nil
}
