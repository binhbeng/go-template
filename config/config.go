package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		AppEnv   string `mapstructure:"app_env" yaml:"app_env"`
		Debug    bool   `mapstructure:"debug" yaml:"debug"`
		EnableBodyLog    bool   `mapstructure:"enable_body_log" yaml:"enable_body_log"`
		Language string `mapstructure:"language" yaml:"language"`
	} `mapstructure:"app" yaml:"app"`

	Jwt struct {
		TTL        int    `mapstructure:"ttl" yaml:"ttl"`
		RefreshTTL int    `mapstructure:"refresh_ttl" yaml:"refresh_ttl"`
		SecretKey  string `mapstructure:"secret_key" yaml:"secret_key"`
	} `mapstructure:"jwt" yaml:"jwt"`

	PostgreDB struct {
		Enable       bool   `mapstructure:"enable" yaml:"enable"`
		Host         string `mapstructure:"host" yaml:"host"`
		Port         int    `mapstructure:"port" yaml:"port"`
		Database     string `mapstructure:"database" yaml:"database"`
		Username     string `mapstructure:"username" yaml:"username"`
		Password     string `mapstructure:"password" yaml:"password"`
		Charset      string `mapstructure:"charset" yaml:"charset"`
		TablePrefix  string `mapstructure:"table_prefix" yaml:"table_prefix"`
		MaxIdleConns int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
		MaxLifetime  string `mapstructure:"max_lifetime" yaml:"max_lifetime"`
		PrintSql     bool   `mapstructure:"print_sql" yaml:"print_sql"`
	} `mapstructure:"postgre" yaml:"postgre"`

	Redis struct {
		Enable   bool   `mapstructure:"enable" yaml:"enable"`
		Host     string `mapstructure:"host" yaml:"host"`
		Port     int    `mapstructure:"port" yaml:"port"`
		Password string `mapstructure:"password" yaml:"password"`
		Database int    `mapstructure:"database" yaml:"database"`
	} `mapstructure:"redis" yaml:"redis"`
}

var Cfg *Config

func init() {
	v := viper.New()
	dir, _ := os.Getwd()

	cfgPath := filepath.Join(dir, "config", "config.yml")
	v.SetConfigFile(cfgPath)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("❌ Error reading config file: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Unable to decode config: %v", err)
	}

	Cfg = &cfg
	log.Println("✅ Config loaded successfully.")
}
