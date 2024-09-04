// usecase/user_usecase.go
package usecase

import (
	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
)

type UserUsecase interface {
	GetUserDetails(id uint) (*models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetUserDetails(id uint) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}
