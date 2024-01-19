package db

import (
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		helper.GetDefaultEnv("DB_HOST", "localhost"),
		helper.GetDefaultEnv("DB_PORT", "5432"),
		helper.GetDefaultEnv("DB_USER", "postgres"),
		helper.GetDefaultEnv("DB_PASS", ""),
		helper.GetDefaultEnv("DB_NAME", "postgres"),
	)

	var newLogger logger.Interface
	if os.Getenv("DB_LOG") == "true" {
		newLogger = logger.Default.LogMode(logger.Info)
	}

	var err error
	db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{
			Logger:      newLogger,
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "YOUR_SCHEMA_NAME.", // schema name
				SingularTable: false,
			},
		},
	)

	if err != nil {
		log.Fatalln("Database connect error", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalln("Database connect error", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maxConn := helper.ParseInt(helper.GetDefaultEnv("DB_CONNECTION_POOL", "0"))
	if maxConn == 0 {
		maxConn = 100
	}
	sqlDB.SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalln("Migrate error", err)
	}

	fmt.Println("> Migration success!!!")
}
