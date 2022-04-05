package product

import (
	"errors"
	"group-project-2/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ============= TESTING ==============

func TestAddProduct(t *testing.T) {
	t.Run("TestNewBookSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		data, err := ProductUseCase.AddProduct(entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10})
		assert.Nil(t, err)
		assert.Equal(t, "Tempe", data.Name)
		assert.Equal(t, uint(3000), data.Price)
		assert.Equal(t, uint(1), data.Seller_ID)
		assert.Equal(t, uint(10), data.Qty)
	})
	t.Run("TestNewBookError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		data, err := ProductUseCase.AddProduct(entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10})
		assert.NotNil(t, err)
		assert.Equal(t, entities.Product{}, data)
	})
}

// ============= SUCCESS MOCK ==============
type mockProductRepository struct{}

func (m mockProductRepository) AddProduct(product entities.Product) (entities.Product, error) {
	return entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10}, nil
}

// ============= ERROR MOCK ==============
type mockProductRepositoryError struct{}

func (m mockProductRepositoryError) AddProduct(product entities.Product) (entities.Product, error) {
	return entities.Product{}, errors.New("add new product error")
}
