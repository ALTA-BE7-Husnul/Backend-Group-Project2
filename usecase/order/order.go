package order

import (
	"group-project-2/delivery/helper"
	_entities "group-project-2/entities"
	_orderRepository "group-project-2/repository/order"
)

type OrderUseCase struct {
	orderRepository _orderRepository.OrderRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository: orderRepo,
	}
}

func (ouc *OrderUseCase) PostOrder(order helper.OrderRequestFormat, idToken int) (_entities.Transaction, error) {
	transaction, err := ouc.orderRepository.PostOrder(order, idToken)
	return transaction, err
}

func (ouc *OrderUseCase) GetOrder(idToken int) ([]_entities.Transaction, error) {
	transaction, err := ouc.orderRepository.GetOrder(idToken)
	return transaction, err
}
