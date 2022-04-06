package product

import (
	"errors"
	_entities "group-project-2/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ============= TESTING ==============

func TestAddProduct(t *testing.T) {
	t.Run("TestAddProductSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		data, err := ProductUseCase.AddProduct(_entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10})
		assert.Nil(t, err)
		assert.Equal(t, "Tempe", data.Name)
		assert.Equal(t, uint(3000), data.Price)
		assert.Equal(t, uint(1), data.Seller_ID)
		assert.Equal(t, uint(10), data.Qty)
	})
	t.Run("TestAddProductError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		data, err := ProductUseCase.AddProduct(_entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Product{}, data)
	})
}
func TestUpdateProductByid(t *testing.T) {
	t.Run("TestUpdateProductByIdSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		data, rows, err := ProductUseCase.UpdateProductById(_entities.Product{Price: 4000}, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, uint(4000), data.Price)
		assert.Equal(t, "Tempe", data.Name)
		assert.Equal(t, 1, rows)
	})
	t.Run("TestUpdateProductByIdError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		data, rows, err := ProductUseCase.UpdateProductById(_entities.Product{Price: 4000}, 1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Product{}, data)
	})
}

func TestDeleteProductById(t *testing.T) {
	t.Run("TestDeleteProductSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		rows, err := ProductUseCase.DeleteProductById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})
	t.Run("TestDeleteProductError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		rows, err := ProductUseCase.DeleteProductById(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

func TestGetAllProduct(t *testing.T) {
	t.Run("TestGetAllProductSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		data, err := ProductUseCase.GetAllProduct()
		assert.Nil(t, err)
		assert.Equal(t, "Tempe", data[0].Name)
		assert.Equal(t, uint(3000), data[0].Price)
	})
	t.Run("TestGetAllProductError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		data, err := ProductUseCase.GetAllProduct()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetProductById(t *testing.T) {
	t.Run("TestGetProductByIdSuccess", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepository{})
		data, rows, err := ProductUseCase.GetProductById(1)
		assert.Nil(t, err)
		assert.Equal(t, "Tempe", data.Name)
		assert.Equal(t, uint(3000), data.Price)
		assert.Equal(t, 1, rows)
	})
	t.Run("TestGetProductByIdError", func(t *testing.T) {
		ProductUseCase := NewProductUseCase(mockProductRepositoryError{})
		data, rows, err := ProductUseCase.GetProductById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Product{}, data)
	})
}

// ============= SUCCESS MOCK ==============
type mockProductRepository struct{}

func (m mockProductRepository) AddProduct(product _entities.Product) (_entities.Product, error) {
	return _entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10}, nil
}
func (m mockProductRepository) GetAllProduct() ([]_entities.Product, error) {
	return []_entities.Product{{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10}}, nil
}

func (m mockProductRepository) GetProductById(id int) (_entities.Product, int, error) {
	return _entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 3000, Qty: 10}, 1, nil
}

func (m mockProductRepository) UpdateProductById(product _entities.Product, id, idToken int) (_entities.Product, int, error) {
	return _entities.Product{Seller_ID: 1, Name: "Tempe", Description: "Makanan Tinggi Protein", Category: "Makanan", Price: 4000, Qty: 10}, 1, nil
}

func (m mockProductRepository) DeleteProductById(id, idToken int) (int, error) {
	return 1, nil
}

// ============= ERROR MOCK ==============
type mockProductRepositoryError struct{}

func (m mockProductRepositoryError) AddProduct(product _entities.Product) (_entities.Product, error) {
	return _entities.Product{}, errors.New("add new product error")
}

func (m mockProductRepositoryError) GetAllProduct() ([]_entities.Product, error) {
	return nil, errors.New("get all error")
}

func (m mockProductRepositoryError) GetProductById(id int) (_entities.Product, int, error) {
	return _entities.Product{}, 0, errors.New("get by id error")
}

func (m mockProductRepositoryError) UpdateProductById(product _entities.Product, id, idToken int) (_entities.Product, int, error) {
	return _entities.Product{}, 0, errors.New("update by id error")
}

func (m mockProductRepositoryError) DeleteProductById(id, idToken int) (int, error) {
	return 0, errors.New("delete error")
}
