package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/user-service/cache"
	"github.com/hariszaki17/library-management/user-service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewUserRepository(db *gorm.DB, cache *cache.Cache) UserRepository {
	return &userRepository{db: db, cache: cache}
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("GetUserByID:%d", id)
	cachedUser, err := r.cache.Get(cacheKey)
	if err == nil && cachedUser != "" {
		// Deserialize cachedUser and return
		user := &models.User{}
		err := json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Cache the result
	// ... (serialize user into a string)
	serializedUser, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedUser), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return &user, nil
}
