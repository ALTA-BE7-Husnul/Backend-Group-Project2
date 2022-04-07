package auth

import (
	"errors"
	_middlewares "group-project-2/delivery/middlewares"
	_entities "group-project-2/entities"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(email string, password string) (string, error) {
	var user _entities.User

	tx := ar.database.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return "failed", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}

	errx := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errx != nil {
		return "filed compare", errx
	}

	token, err := _middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return "create token failed", err
	}

	return token, nil

}
