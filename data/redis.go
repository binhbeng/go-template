package data

import (
	"context"
	"fmt"
	"log"

	"github.com/binhbeng/goex/config"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func initRedis() {

	addr := fmt.Sprintf("%s:%d", config.C.Redis.Host, config.C.Redis.Port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.Database,
	})
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis connection failed：" + err.Error())
	}
	log.Println("✅ Redis connected")
}
