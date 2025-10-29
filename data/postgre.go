package data

import (
	"fmt"
	"time"
	// c "github.com/wannanbigpig/gin-layout/config"
	// log "github.com/wannanbigpig/gin-layout/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var PostgreDB *gorm.DB

type Writer interface {
	Printf(string, ...interface{})
}

type WriterLog struct{}

func (w WriterLog) Printf(format string, args ...interface{}) {
	// if c.Config.Mysql.PrintSql {
	// 	log.Logger.Sugar().Infof(format, args...)
	// }
}

func initPostgre() {
	logConfig := logger.New(
		WriterLog{},
		logger.Config{
			SlowThreshold:             0,                                        // 慢 SQL 阈值
			LogLevel:                  logger.LogLevel(4), // 日志级别 c.Config.Mysql.LogLevel
			IgnoreRecordNotFoundError: false,                                    // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                                    // 是否启用彩色打印
		},
	)

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表名前缀 c.Config.Mysql.TablePrefix
			// SingularTable: true,  // 使用单数表名
		},
		Logger: logConfig,
	}

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
	// 	c.Config.Mysql.Username,
	// 	c.Config.Mysql.Password,
	// 	c.Config.Mysql.Host,
	// 	c.Config.Mysql.Port,
	// 	c.Config.Mysql.Database,
	// 	c.Config.Mysql.Charset,
	// )
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
	// 	"binhbeng",
	// 	"1998",
	// 	"localhost",
	// 	5432,
	// 	"goapp",
	// 	"utf8",
	// )
 	dsn := "host=localhost user=binhbeng password=1998 dbname=goapp port=5432 sslmode=disable"
	var err error
	
	PostgreDB, err = gorm.Open(postgres.Open(dsn), configs)

	if err != nil {
		panic("PostgreSQL connection failed：" + err.Error())
	} else {
		fmt.Println("🔍 ~ initPostgre ~ data/postgre.go:62 ~ dsn:", dsn);
	}


	sqlDB, _ := PostgreDB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
