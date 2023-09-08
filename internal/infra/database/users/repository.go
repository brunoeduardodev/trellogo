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

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (r *UserRepository) Create(input CreateUserInput) error {
	user := UserModel{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	result := r.db.Create(&user)

	return result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}
