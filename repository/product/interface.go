package product

import _entities "group-project-2/entities"

type ProductRepositoryInterface interface {
	AddProduct(product _entities.Product) (_entities.Product, error)
}
