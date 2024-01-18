package cache

import (
	"TQP0403/todo-list/src/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICacheService interface {
	HSet(key string, val interface{}, ttl time.Duration) error
	HGet(key string, val interface{}) error
	HDel(key string) error

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

func (s *CacheService) HSet(key string, val interface{}, ttl time.Duration) error {
	if ttl == 0 {
		ttl = s.ttlDefault
	}

	return s.client.HSet(s.ctx, key, val, ttl).Err()
}

func (s *CacheService) HGet(key string, val interface{}) error {
	return s.client.HGetAll(s.ctx, key).Scan(&val)
}

func (s *CacheService) HDel(key string) error {
	return s.client.HDel(s.ctx, key).Err()
}

func (s *CacheService) Set(key string, val interface{}, ttl time.Duration) error {
	if ttl == 0 {
		ttl = s.ttlDefault
	}

	return s.client.Set(s.ctx, key, val, ttl).Err()
}

func (s *CacheService) Get(key string, val interface{}) error {
	return s.client.Get(s.ctx, key).Scan(&val)
}

func (s *CacheService) Del(key string) error {
	return s.client.Del(s.ctx, key).Err()
}
