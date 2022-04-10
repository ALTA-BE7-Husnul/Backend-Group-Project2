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

func (ar *AuthRepository) Login(email string, password string) (string, uint, error) {
	var user _entities.User

	tx := ar.database.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return "failed", 0, tx.Error
	}

	if tx.RowsAffected == 0 {
		return "user not found", 0, errors.New("user not found")
	}

	errx := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errx != nil {
		return "filed compare", 0, errx
	}

	token, err := _middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return "create token failed", 0, err
	}
	id := user.ID
	return token, id, nil

}
