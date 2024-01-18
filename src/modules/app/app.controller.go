package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	service IAppService
}

func NewController(service IAppService) *AppController {
	return &AppController{service: service}
}

func (ctrl *AppController) Register(router *gin.Engine) {
	group := router.Group("/")
	{
		group.GET("/", ctrl.handleHello)
		group.GET("/ping", ctrl.handlePing)
	}
}

func (ctrl *AppController) handleHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctrl.service.HandleHello())
}

func (ctrl *AppController) handlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctrl.service.HandlePing())
}
