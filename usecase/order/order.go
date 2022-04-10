package order

import (
	_entities "group-project-2/entities"
	_cartRepository "group-project-2/repository/cart"
	_orderRepository "group-project-2/repository/order"
	_productRepository "group-project-2/repository/product"
)

type OrderUseCase struct {
	orderRepository   _orderRepository.OrderRepositoryInterface
	cartRepository    _cartRepository.CartRepositoryInterface
	productRepository _productRepository.ProductRepositoryInterface
}

func NewOrderUseCase(orderRepo _orderRepository.OrderRepositoryInterface, cartRepo _cartRepository.CartRepositoryInterface, productRepo _productRepository.ProductRepositoryInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		orderRepository:   orderRepo,
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (ouc *OrderUseCase) PostOrder(order _entities.Transaction, orderCartID []uint, idToken int) (_entities.Transaction, int, error) {
	carts, getErr := ouc.cartRepository.GetAll()
	if getErr != nil {
		return order, 0, getErr
	}
	for i := range orderCartID {
		order.Total += carts[i].Total
	}
	order.Status = "paid"
	transaction, rows, err := ouc.orderRepository.PostOrder(order, orderCartID)
	return transaction, rows, err
}

func (ouc *OrderUseCase) GetOrder(idToken int) ([]_entities.Transaction, error) {
	transaction, err := ouc.orderRepository.GetOrder(idToken)
	return transaction, err
}
