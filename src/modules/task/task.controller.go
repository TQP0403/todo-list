package task

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/modules/jwt"
	"TQP0403/todo-list/src/modules/task/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service    ITaskService
	jwtService jwt.IJwtService
}

func NewController(service *TaskService, jwtService *jwt.JwtService) *TaskController {
	return &TaskController{service: service, jwtService: jwtService}
}

func (ctrl *TaskController) Register(router *gin.Engine) {

	group := router.Group("/api/task")
	{
		group.POST("/", ctrl.jwtService.JwtMiddleware(), ctrl.handleCreateTask)
		group.GET("/", ctrl.jwtService.JwtMiddleware(), ctrl.handleGetListTask)
		group.GET("/:id", ctrl.jwtService.JwtMiddleware(), ctrl.handleGetTaskById)
		group.PUT("/:id", ctrl.jwtService.JwtMiddleware(), ctrl.handleUpdateTask)
		group.DELETE("/:id", ctrl.jwtService.JwtMiddleware(), ctrl.handleDeleteTask)
	}
}

func (ctrl *TaskController) handleCreateTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	var reqData dtos.CreateTaskDto
	if err := ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, cusErr.GetErrorResponse())
		return
	}
	reqData.UserID = userId

	if _, err := ctrl.service.CreateTask(&reqData); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, cusErr.GetErrorResponse())
		return
	}

	ctx.JSON(http.StatusOK, common.NewSimpleResponse())
}

func (ctrl *TaskController) handleGetListTask(ctx *gin.Context) {
	pQuery := common.NewPagination()
	userId := jwt.GetUserId(ctx)

	if err := ctx.ShouldBindQuery(pQuery); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr.GetErrorResponse())
		return
	}

	if tasks, err := ctrl.service.GetListTask(userId, pQuery); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr.GetErrorResponse())
	} else {
		res := common.NewPaginationResponse(pQuery, tasks)
		ctx.JSON(http.StatusOK, res.GetSuccessResponse())
	}
}

func (ctrl *TaskController) handleGetTaskById(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	id := helper.ParseInt(ctx.Param("id"))

	if task, err := ctrl.service.GetTaskById(userId, id); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr.GetErrorResponse())
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(task))
	}
}

func (ctrl *TaskController) handleUpdateTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)
	var reqData dtos.UpdateTaskDto

	if err := ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, cusErr.GetErrorResponse())
		return
	}

	reqData.ID = helper.ParseInt(ctx.Param("id"))
	if err := ctrl.service.UpdateTask(userId, &reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, cusErr.GetErrorResponse())
	} else {
		ctx.JSON(http.StatusOK, common.NewSimpleResponse())
	}
}

func (ctrl *TaskController) handleDeleteTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)
	id := helper.ParseInt(ctx.Param("id"))

	if err := ctrl.service.DeleteTask(userId, id); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, cusErr.GetErrorResponse())
		return
	}

	ctx.JSON(http.StatusOK, common.NewSimpleResponse())
}
