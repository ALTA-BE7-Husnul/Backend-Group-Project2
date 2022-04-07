package cart

import (
	"fmt"
	_entities "group-project-2/entities"
	_cartRepository "group-project-2/repository/cart"
)

type CartUseCase struct {
	cartRepository _cartRepository.CartRepositoryInterface
}

func NewCartUseCase(cartRepo _cartRepository.CartRepositoryInterface) CartUseCaseInterface {
	return &CartUseCase{
		cartRepository: cartRepo,
	}
}

func (uuc *CartUseCase) PostCart(cart _entities.Cart, idToken int) (_entities.Cart, int, int, error) {
	fmt.Println("ini di usecase ", cart)
	cart, rows, err := uuc.cartRepository.PostCart(cart, idToken)
	return cart, idToken, rows, err
}
