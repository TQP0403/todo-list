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
		group.POST("/", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleCreateTask)
		group.GET("/", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleGetListTask)
		group.GET("/:id", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleGetTaskById)
		group.PUT("/:id", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleUpdateTask)
		group.DELETE("/:id", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleDeleteTask)
	}
}

func (ctrl *TaskController) handleCreateTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	var reqData dtos.CreateTaskDto
	if err := ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
		return
	}
	reqData.UserID = userId

	if _, err := ctrl.service.CreateTask(&reqData); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
		return
	}

	ctx.JSON(http.StatusOK, common.NewSimpleResponse())
}

func (ctrl *TaskController) handleGetListTask(ctx *gin.Context) {
	pQuery := common.BindPagination(ctx)
	userId := jwt.GetUserId(ctx)

	tasks, err := ctrl.service.GetListTask(userId, &pQuery)

	if err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
		return
	}

	res := &common.PaginationResponse{
		Metadata: pQuery,
		Rows:     tasks,
	}

	ctx.JSON(http.StatusOK, common.NewListResponse(res))
}

func (ctrl *TaskController) handleGetTaskById(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	id := helper.ParseInt(ctx.Param("id"))

	if task, err := ctrl.service.GetTaskById(userId, id); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(task))
	}
}

func (ctrl *TaskController) handleUpdateTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	var reqData dtos.UpdateTaskDto

	if err := ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
		return
	}

	reqData.ID = helper.ParseInt(ctx.Param("id"))
	if err := ctrl.service.UpdateTask(userId, &reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
	} else {
		ctx.JSON(http.StatusOK, common.NewSimpleResponse())
	}
}

func (ctrl *TaskController) handleDeleteTask(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)
	id := helper.ParseInt(ctx.Param("id"))

	if err := ctrl.service.DeleteTask(userId, id); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
		return
	}

	ctx.JSON(http.StatusOK, common.NewSimpleResponse())
}
