package server

import (
	"TQP0403/todo-list/src/db"
	"TQP0403/todo-list/src/modules/app"
	"TQP0403/todo-list/src/modules/auth"
	"TQP0403/todo-list/src/modules/cache"
	"TQP0403/todo-list/src/modules/file"
	"TQP0403/todo-list/src/modules/jwt"
	"TQP0403/todo-list/src/modules/task"
	"os"

	"github.com/gin-gonic/gin"
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

func Default(myDb db.IMyGormService) *Router {
	// repos
	authRepo := auth.NewRepo(myDb.GetDB())
	taskRepo := task.NewRepo(myDb.GetDB())

	// services
	cacheService := cache.NewDefaultCacheService()
	appService := app.NewService()
	jwtService := jwt.NewJwtService(os.Getenv("JWT_SECRET"))
	authService := auth.NewService(authRepo, jwtService, cacheService)
	taskService := task.NewService(taskRepo, cacheService)
	fileService := file.NewService()

	// controllers
	appController := app.NewController(appService)
	authController := auth.NewController(authService, jwtService)
	taskController := task.NewController(taskService, jwtService)
	fileController := file.NewController(fileService)

	return &Router{
		controllers: []IController{
			appController,
			authController,
			taskController,
			fileController,
		},
	}
}
