package cart

import (
	_entities "group-project-2/entities"
)

type CartRepositoryInterface interface {
	PostCart(cart _entities.Cart, idToken int) (_entities.Cart, error)
}
