package database

import (
	"log"
	"os"
	"time"

	"github.com/goer-project/goer/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbConfig gorm.Dialector
	DB       *gorm.DB
)

// Gorm init gorm
func Gorm() *gorm.DB {
	switch config.NewConfig.Database.Connection {
	case "mysql":
		GormMysqlConfig()
	case "sqlite":
		GormSqliteConfig()
	default:
		GormMysqlConfig()
	}

	// Custom logger
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open DB
	db, err := gorm.Open(dbConfig, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   customLogger,
	})
	if err != nil {
		log.Println(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.NewConfig.Database.Mysql.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.NewConfig.Database.Mysql.MaxOpenConnection)

	return db
}

func GormMysqlConfig() {
	dbConfig = mysql.New(mysql.Config{
		DSN: config.NewConfig.Database.Mysql.Dsn(),
	})
}

func GormSqliteConfig() {
	dbConfig = sqlite.Open(config.NewConfig.Database.Sqlite.Database)
}
