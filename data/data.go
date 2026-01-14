package data

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var once sync.Once
var RedisDB *redis.Client
var PostgreDB *gorm.DB

func InitData() {
	once.Do(func() {
		initPostgre()
		initRedis()
	})
}

func initPostgre() {
	db, err := NewPostgreDB()
	if err != nil {
		panic(err)
	}
	PostgreDB = db
}

func initRedis() {
	db, err := NewRedis()
	if err != nil {
		panic(err)
	}
	RedisDB = db
}
