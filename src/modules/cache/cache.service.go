package cache

import (
	"TQP0403/todo-list/src/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheService interface {
	Set(key string, val interface{}, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
}

type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewDefaultCacheService() *CacheService {
	redisStr := "redis"
	if config.Getenv("REDIS_TLS", "false") == "true" {
		redisStr = "rediss"
	}

	redisUrl := fmt.Sprintf("%s://%s:%s@%s:%s/0",
		redisStr,
		config.Getenv("REDIS_USER", ""),
		config.Getenv("REDIS_PASS", ""),
		config.Getenv("REDIS_HOST", "localhost"),
		config.Getenv("REDIS_PORT", "6379"),
	)
	if opt, err := redis.ParseURL(redisUrl); err != nil {
		panic(err)
	} else {
		return &CacheService{ctx: context.Background(), client: redis.NewClient(opt)}
	}
}

func (s *CacheService) Set(key string, val interface{}, ttl time.Duration) error {
	if ttl == 0 {
		ttl = time.Hour
	}

	return s.client.Set(s.ctx, key, val, ttl).Err()
}

func (s *CacheService) Get(key string) (string, error) {
	return s.client.Get(s.ctx, key).Result()
}

func (s *CacheService) Del(key string) error {
	return s.client.Del(s.ctx, key).Err()
}
