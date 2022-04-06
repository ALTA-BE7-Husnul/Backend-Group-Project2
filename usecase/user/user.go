package user

import (
	_entities "group-project-2/entities"
	_userRepository "group-project-2/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) PostUser(user _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.PostUser(user)
	return users, err
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetUser(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUser(id)
	return user, rows, err
}
func (uuc *UserUseCase) DeleteUser(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.DeleteUser(id)
	if user.Name == "" {
		return user, rows, err
	}
	return user, rows, nil
}

func (uuc *UserUseCase) PutUser(id int, user _entities.User) (_entities.User, error) {
	user.ID = uint(id)
	_, _, err := uuc.userRepository.GetUser(id)
	if err != nil {
		return user, err
	}

	userPut, errPut := uuc.userRepository.PutUser(user)
	return userPut, errPut
}
