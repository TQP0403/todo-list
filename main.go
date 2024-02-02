package main

import (
	"TQP0403/todo-list/src/config"
	"TQP0403/todo-list/src/db"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/middlewares"
	"TQP0403/todo-list/src/server"
	"fmt"
	"os"

	docs "TQP0403/todo-list/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a todo-list backend server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupRouter(myDb db.IMyGormService) *gin.Engine {
	r := gin.Default()

	r.ForwardedByClientIP = true
	r.Use(middlewares.CORSMiddleware())

	serv := server.Default(myDb.GetDB())
	serv.Register(r)

	return r
}

func main() {
	// load .env file
	config.Init()

	// connect database
	myDb := db.Init()

	// setup gin router
	r := setupRouter(myDb)

	if env := os.Getenv("GIN_ENV"); env != "production" {
		// run auto migration with goroutine
		go myDb.Migrate()
		// swagger
		setupSwagger(r)
	}

	adrr := fmt.Sprintf(":%s", helper.GetDefaultEnv("GIN_PORT", "8080"))

	// run gin app
	r.Run(adrr)
}
