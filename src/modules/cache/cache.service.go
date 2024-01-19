package cache

import (
	"TQP0403/todo-list/src/helper"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheService interface {
	Set(key string, val interface{}, ttl time.Duration) error
	Get(key string, val interface{}) error
	Del(key string) error
}

type CacheService struct {
	client     *redis.Client
	ctx        context.Context
	ttlDefault time.Duration
}

func NewDefaultCacheService() *CacheService {
	redisStr := "redis"
	if os.Getenv("REDIS_TLS") == "true" {
		redisStr = "rediss"
	}

	redisUrl := fmt.Sprintf("%s://%s:%s@%s:%s/0",
		redisStr,
		helper.GetDefaultEnv("REDIS_USER", ""),
		helper.GetDefaultEnv("REDIS_PASS", ""),
		helper.GetDefaultEnv("REDIS_HOST", "localhost"),
		helper.GetDefaultEnv("REDIS_PORT", "6379"),
	)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	return &CacheService{
		ctx:        context.Background(),
		client:     client,
		ttlDefault: time.Hour,
	}
}

func (s *CacheService) Set(key string, val interface{}, ttl time.Duration) error {
	return s.client.Set(s.ctx, key, val, helper.GetDefaultNumber(ttl, s.ttlDefault)).Err()
}

func (s *CacheService) Get(key string, val interface{}) error {
	return s.client.Get(s.ctx, key).Scan(&val)
}

func (s *CacheService) Del(key string) error {
	return s.client.Del(s.ctx, key).Err()
}
