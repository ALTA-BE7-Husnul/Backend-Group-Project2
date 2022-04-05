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
