package config

import (
	"fmt"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

func RedisClient(cfg Config) *cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.REDIS_HOST, cfg.REDIS_PORT),
		Password: cfg.REDIS_PASS,
		DB:       0,
	})

	client := cache.New(&cache.Options{
		Redis: rdb,
	})

	return client
}
