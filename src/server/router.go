package server

import (
	"TQP0403/todo-list/src/modules/app"
	"TQP0403/todo-list/src/modules/auth"
	"TQP0403/todo-list/src/modules/cache"
	"TQP0403/todo-list/src/modules/jwt"
	"TQP0403/todo-list/src/modules/task"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IController interface {
	Register(router *gin.Engine)
}

type Router struct {
	controllers []IController
}

func (r *Router) Register(router *gin.Engine) {
	// controllers register
	for _, ctrl := range r.controllers {
		ctrl.Register(router)
	}
}

func Default(db *gorm.DB) *Router {
	// secret := []byte(os.Getenv("JWT_SECRET"))
	// repos
	authRepo := auth.NewRepo(db)
	taskRepo := task.NewRepo(db)

	// services
	cacheService := cache.NewDefaultCacheService()
	appService := app.NewService()
	jwtService := jwt.NewJwtService([]byte(os.Getenv("JWT_SECRET")))
	authService := auth.NewService(authRepo, jwtService, cacheService)
	taskService := task.NewService(taskRepo)

	// controllers
	appController := app.NewController(appService)
	authController := auth.NewController(authService)
	taskController := task.NewController(taskService, jwtService)

	return &Router{
		controllers: []IController{
			appController,
			authController,
			taskController,
		},
	}
}
