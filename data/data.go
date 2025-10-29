package data

import (
	// c "github.com/wannanbigpig/gin-layout/config"
	"sync"
)

var once sync.Once

func InitData() {
	once.Do(func() {
		if true {
			initPostgre()
			initRedis()
		}

		// if c.Config.Redis.Enable {
		// 	// 初始化 redis
		// 	initRedis()
		// }
	})
}
