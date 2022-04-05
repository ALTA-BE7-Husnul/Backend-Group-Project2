package user

import "gorm.io/gorm"

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetById(id int) (_entities.User, int, error) {
	// write your code here
	return user, int(x), nil
}
