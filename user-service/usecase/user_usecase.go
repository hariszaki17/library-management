// usecase/user_usecase.go
package usecase

import (
	"context"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hariszaki17/library-management/user-service/config"
	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
	"gorm.io/gorm"
)

var jwtSecret = []byte(config.Data.SecretJWT)

type UserUsecase interface {
	GetUserDetails(ctx context.Context, id uint) (*models.User, error)
	Authenticate(ctx context.Context, username, password string) (*models.AuthUseCaseResp, error)
	VerifyJWT(ctx context.Context, token string) (*models.User, error)
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

func (u *userUsecase) Authenticate(ctx context.Context, username, password string) (*models.AuthUseCaseResp, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	// TODO: Need to be hashed
	if username == user.Username && password == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return nil, err
		}
		return &models.AuthUseCaseResp{
			User:  *user,
			Token: tokenString,
		}, nil
	}

	return nil, errors.New("invalid username or password")
}

func (u *userUsecase) VerifyJWT(ctx context.Context, tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &models.User{}
		valID, ok := claims["id"]
		if ok {
			user.ID = uint(valID.(float64))
		}

		valUsername, ok := claims["username"]
		if ok {
			user.Username = valUsername.(string)
		}
		return user, nil
	}

	return nil, errors.New("invalid token")
}
