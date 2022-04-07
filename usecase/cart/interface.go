package cart

import (
	_entities "group-project-2/entities"
)

type CartUseCaseInterface interface {
	PostCart(cart _entities.Cart, idToken int) (_entities.Cart, int, int, error)
	GetAll() ([]_entities.Cart, error)
}