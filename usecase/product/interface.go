package product

import _entities "group-project-2/entities"

type ProductUseCaseInterface interface {
	AddProduct(product _entities.Product) (_entities.Product, error)
}
