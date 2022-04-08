package order

import (
	"group-project-2/delivery/helper"
	_entities "group-project-2/entities"
)

type OrderUseCaseInterface interface {
	PostOrder(order helper.OrderRequestFormat, idToken int) (_entities.Transaction, error)
	GetOrder(idToken int) ([]_entities.Transaction, error)
}
