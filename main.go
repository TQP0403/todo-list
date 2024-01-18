package main

import (
	"TQP0403/todo-list/src/config"
	"TQP0403/todo-list/src/db"
	"TQP0403/todo-list/src/middlewares"
	"TQP0403/todo-list/src/server"
	"fmt"
	"log"
	"os"

	_ "TQP0403/todo-list/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server todo-list api.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/
func main() {
	config.Init()

	db.Init()

	router := gin.Default()
	router.ForwardedByClientIP = true

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())

	s := server.Default(db.GetDB())
	s.Register(router)

	env := os.Getenv("ENV")
	if env != "production" {
		db.Migrate()
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalf("Go Gin run err: %s\n", err)
	}
}
