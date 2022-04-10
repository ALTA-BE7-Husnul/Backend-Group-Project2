package order

import (
	_entities "group-project-2/entities"
)

type OrderRepositoryInterface interface {
	PostOrder(order _entities.Transaction, orderCartID []uint) (_entities.Transaction, int, error)
	GetOrder(idToken int) ([]_entities.Transaction, error)
}
