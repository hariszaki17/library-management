// usecase/user_usecase.go
package usecase

import (
	"context"
	"errors"

	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUserDetails(ctx context.Context, id uint) (*models.User, error)
	Authenticate(ctx context.Context, username, password string) (*models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetUserDetails(ctx context.Context, id uint) (*models.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *userUsecase) Authenticate(ctx context.Context, username, password string) (*models.User, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}
	if username == user.Username && password == user.Password {
		return &models.User{
			Model: gorm.Model{
				ID: user.ID,
			},
			Username: username,
		}, nil
	}

	return nil, errors.New("invalid username or password")
}
