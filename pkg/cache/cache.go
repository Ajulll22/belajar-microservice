package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

func RedisClient(REDIS_HOST string, REDIS_PORT string, REDIS_PASS string) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
		Password: REDIS_PASS,
		DB:       0,
	})

	client := cache.New(&cache.Options{
		Redis: rdb,
	})

	return &redisCache{
		cache: client,
	}
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, value interface{}) error
}

type redisCache struct {
	cache *cache.Cache
}

func (c *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	i := cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
	}

	return c.cache.Set(&i)
}

func (c *redisCache) Get(ctx context.Context, key string, value interface{}) error {
	err := c.cache.Get(ctx, key, value)
	if err == cache.ErrCacheMiss {
		return nil
	}

	return err
}

func GetCacheKey(key string, additionalKey ...int) string {
	for _, id := range additionalKey {
		key += fmt.Sprintf("_%d", id)
	}
	return key
}
