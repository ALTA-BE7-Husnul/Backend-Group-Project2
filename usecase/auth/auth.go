package auth

import (
	_authRepository "group-project-2/repository/auth"
)

type AuthUseCase struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo _authRepository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}

func (auc *AuthUseCase) Login(email string, password string) (string, uint, error) {
	token, id, err := auc.authRepository.Login(email, password)
	return token, id, err
}
