package user

import (
	_entities "group-project-2/entities"
)

type UserRepositoryInterface interface {
	PostUser(user _entities.User) (_entities.User, error)
	GetAll() ([]_entities.User, error)
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, int, error)
	PutUser(user _entities.User, id int) (_entities.User, error)
}
