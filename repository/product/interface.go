package product

import _entities "group-project-2/entities"

type ProductRepositoryInterface interface {
	AddProduct(product _entities.Product) (_entities.Product, error)
	UpdateProductById(product _entities.Product, id, idToken int) (_entities.Product, int, error)
	DeleteProductById(id, idToken int) (int, error)
	GetAllProduct() ([]_entities.Product, error)
	GetProductById(id int) (_entities.Product, int, error)
}
