// cache/cache.go
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
	ctx    context.Context
}

func NewCache(redisAddr string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &Cache{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (c *Cache) Get(key string) (string, error) {
	val, err := c.client.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Cache miss
	}
	return val, err
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(c.ctx, key, value, ttl).Err()
}
