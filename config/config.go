package config

import (
	"log"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var (
	Cfg  *Config
	once sync.Once
)

type Config struct {
	App struct {
		AppEnv        string `env:"APP_ENV"`
		Debug         bool   `env:"DEBUG"`
		EnableBodyLog bool   `env:"ENABLE_BODY_LOG"`
		Language      string `env:"LANGUAGE"`
		Socket        bool   `env:"SOCKET"`
	}

	Jwt struct {
		TTL        int    `env:"JWT_TTL"`
		RefreshTTL int    `env:"JWT_REFRESH_TTL"`
		SecretKey  string `env:"JWT_SECRET_KEY"`
	}

	PostgreDB struct {
		Enable       bool   `env:"POSTGRE_ENABLE"`
		Host         string `env:"POSTGRE_HOST"`
		Port         int    `env:"POSTGRE_PORT"`
		Database     string `env:"POSTGRE_DATABASE"`
		Username     string `env:"POSTGRE_USERNAME"`
		Password     string `env:"POSTGRE_PASSWORD"`
		Charset      string `env:"POSTGRE_CHARSET"`
		TablePrefix  string `env:"POSTGRE_TABLE_PREFIX"`
		MaxIdleConns int    `env:"POSTGRE_MAX_IDLE_CONNS"`
		MaxOpenConns int    `env:"POSTGRE_MAX_OPEN_CONNS"`
		MaxLifetime  string `env:"POSTGRE_MAX_LIFETIME"`
		PrintSql     bool   `env:"POSTGRE_PRINT_SQL"`
	}

	Redis struct {
		Enable   bool   `env:"REDIS_ENABLE"`
		Host     string `env:"REDIS_HOST"`
		Port     int    `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
		Database int    `env:"REDIS_DATABASE"`
	}
}

func Load() {
	once.Do(func() {
		environment := os.Getenv("APP_ENV")
		if environment == "local" || environment == "dev" || environment == "" {
			_ = godotenv.Load()
		}

		var c Config
		if err := env.Parse(&c); err != nil {
			log.Fatalf("❌ Load config failed: %v", err)
		}

		Cfg = &c
		log.Println("✅ Config loaded successfully.")
	})
}
