package user

import (
	_entities "group-project-2/entities"
)

type UserUseCaseInterface interface {
	PostUser(user _entities.User) (_entities.User, error)
	GetAll() ([]_entities.User, error)
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, int, error)
	PutUser(id int, user _entities.User) (_entities.User, error)
}
