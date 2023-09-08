package users

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) List() (*[]UserModel, error) {
	var users []UserModel
	result := r.db.Find(&users)

	return &users, result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}
