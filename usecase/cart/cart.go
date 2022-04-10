package cart

import (
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
	cart, rows, err := uuc.cartRepository.PostCart(cart, idToken)
	return cart, idToken, rows, err
}

func (uuc *CartUseCase) GetAll() ([]_entities.Cart, error) {
	cart, err := uuc.cartRepository.GetAll()
	return cart, err
}
func (uuc *CartUseCase) PutCart(cart _entities.Cart, idToken int) (_entities.Cart, error) {
	cart, err := uuc.cartRepository.PutCart(cart, idToken)
	return cart, err
}
func (uuc *CartUseCase) DeleteCart(id int) (_entities.Cart, int, error) {
	cart, rows, err := uuc.cartRepository.DeleteCart(id)
	if cart.Product_ID == 0 {
		return cart, rows, err
	}
	return cart, rows, nil
}
