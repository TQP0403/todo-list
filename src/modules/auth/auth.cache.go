package auth

import (
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/cache"
	"fmt"
	"time"
)

type IAuthCache interface {
	GetCacheUser(id int) (*models.User, error)
	SetCacheUser(user *models.User) error
	DelCacheUser(id int) error
}

type AuthCache struct {
	cacheService cache.ICacheService
}

func NewAuthCache(cacheService *cache.CacheService) *AuthCache {
	return &AuthCache{cacheService: cacheService}
}

func (service *AuthCache) GetCacheUser(id int) (*models.User, error) {
	user := &models.User{}
	key := fmt.Sprintf("user:%d", id)
	if err := service.cacheService.Get(key, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (service *AuthCache) SetCacheUser(user *models.User) error {
	ttl := time.Minute * 5
	key := fmt.Sprintf("user:%d", user.ID)
	if err := service.cacheService.Set(key, user, ttl); err != nil {
		return err
	}

	return nil
}

func (service *AuthCache) DelCacheUser(id int) error {
	key := fmt.Sprintf("user:%d", id)
	if err := service.cacheService.Del(key); err != nil {
		return err
	}

	return nil
}
