package user

import (
	_entities "group-project-2/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) PostUser(user _entities.User) (_entities.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return _entities.User{}, err
	}

	user.Password = string(hashedPassword)

	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetUser(id int) (_entities.User, int, error) {
	var user _entities.User
	tx := ur.database.Find(&user, id)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, nil
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) DeleteUser(id int) (_entities.User, int, error) {
	var user _entities.User

	tx := ur.database.Where("ID = ?", id).Delete(&user)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, nil
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) PutUser(user _entities.User) (_entities.User, error) {

	userSave := user

	tx := ur.database.Save(&userSave)
	if tx != nil {
		return user, tx.Error
	}
	return user, nil
}
