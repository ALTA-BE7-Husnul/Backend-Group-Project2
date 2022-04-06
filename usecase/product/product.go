package product

import (
	_entities "group-project-2/entities"
	_productRepository "group-project-2/repository/product"
)

type ProductUseCase struct {
	productRepository _productRepository.ProductRepositoryInterface
}

func NewProductUseCase(productRepo _productRepository.ProductRepositoryInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		productRepository: productRepo,
	}
}

func (puc *ProductUseCase) AddProduct(product _entities.Product) (_entities.Product, error) {
	product, err := puc.productRepository.AddProduct(product)
	return product, err
}

func (puc *ProductUseCase) UpdateProductById(product _entities.Product, id, idToken int) (_entities.Product, int, error) {
	data, rows, err := puc.productRepository.UpdateProductById(product, id, idToken)
	return data, rows, err
}

func (puc *ProductUseCase) DeleteProductById(id, idToken int) (int, error) {
	rows, err := puc.productRepository.DeleteProductById(id, idToken)
	return rows, err
}

func (puc *ProductUseCase) GetAllProduct() ([]_entities.Product, error) {
	products, err := puc.productRepository.GetAllProduct()
	return products, err
}

func (puc *ProductUseCase) GetProductById(id int) (_entities.Product, int, error) {
	product, rows, err := puc.productRepository.GetProductById(id)
	return product, rows, err
}
