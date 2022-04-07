package cart

import (
	_entities "group-project-2/entities"
)

type CartRepositoryInterface interface {
	PostCart(cart _entities.Cart, idToken int) (_entities.Cart, int, error)
	GetAll() ([]_entities.Cart, error)
	PutCart(cart _entities.Cart, idToken int) (_entities.Cart, error)
	DeleteCart(id int) (_entities.Cart, int, error)
}
