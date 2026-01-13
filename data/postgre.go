package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/binhbeng/goex/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var PgxDB *pgxpool.Pool

type Writer interface {
	Printf(string, ...any)
}

type WriterLog struct{}

func (w WriterLog) Printf(format string, args ...any) {
	if config.Cfg.PostgreDB.PrintSql {
		s := fmt.Sprintf(format, args...)
		fmt.Println(s)
	}
}

func NewPostgreDB() (*gorm.DB, error) {
	logConfig := logger.New(
		WriterLog{},
		logger.Config{
			SlowThreshold:             0,
			LogLevel:                  logger.LogLevel(4),
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",
		},
		Logger: logConfig,
	}

	dbCfg := config.Cfg.PostgreDB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.Database, dbCfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), configs)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQL connection failed: %w", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ PostgreSQL connection success")
	return db, nil
}

func NewPgxDB() (*pgxpool.Pool, error) {
	dbCfg := config.Cfg.PostgreDB
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	config.MaxConns = 10
	config.MinConns = 0
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("✅ Pgx connection success")
	return pool, nil
}
