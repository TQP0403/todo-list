package db

import (
	"TQP0403/todo-list/src/common"
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

type MyGormService struct {
	db *gorm.DB
}

type IMyGormService interface {
	GetDB() *gorm.DB
	Migrate()
}

func Init() *MyGormService {
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

	db, err := gorm.Open(
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

	return &MyGormService{db: db}
}

func (server *MyGormService) GetDB() *gorm.DB {
	return server.db
}

func (server *MyGormService) Migrate() {
	var m []interface{}
	m = append(m, &models.User{}, &models.Task{})

	if err := common.IsModel(m...); err != nil {
		log.Fatalln("Check model error:", err)
	}

	if err := server.db.AutoMigrate(m...); err != nil {
		log.Fatalln("Migrate error:", err)
	}

	log.Println("Migration success!!!")
}
