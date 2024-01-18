package db

import (
	"TQP0403/todo-list/src/config"
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
		config.Getenv("DB_HOST", "localhost"),
		config.Getenv("DB_PORT", "5432"),
		config.Getenv("DB_USER", "postgres"),
		config.Getenv("DB_PASS", ""),
		config.Getenv("DB_NAME", "postgres"),
	)

	var newLogger logger.Interface
	if dbLog := config.Getenv("DB_LOG", "false"); dbLog == "true" {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		)
	}

	var err error
	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "YOUR_SCHEMA_NAME.", // schema name
				SingularTable: false,
			},
		},
	)

	if err != nil {
		log.Fatalln("Database connect error", err)
	}
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
