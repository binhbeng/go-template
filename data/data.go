package data

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var RedisDB *redis.Client

func InitData() {
	once.Do(func() {
		initPgx()
		initRedis()
	})
}

func initPgx() {
	db, err := NewPgxDB()
	if err != nil {
		panic(err)
	}
	PgxDB = db
}

func initRedis() {
	db, err := NewRedis()
	if err != nil {
		panic(err)
	}
	RedisDB = db
}
