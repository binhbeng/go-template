package data

import (
	"context"
	"fmt"
	"log"

	"github.com/binhbeng/goex/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.Database,
	})

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	log.Println("✅ Redis connected")
	return rdb, nil
}