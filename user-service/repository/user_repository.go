package repository

import (
	"github.com/hariszaki17/library-management/user-service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
