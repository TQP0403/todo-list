package file

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/modules/file/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type FileController struct {
	service IFileService
}

func NewController(service *FileService) *FileController {
	return &FileController{service: service}
}

func (ctrl *FileController) Register(router *gin.Engine) {
	group := router.Group("/api/file")
	{
		group.POST("/upload", ctrl.handleUploadFile)
	}
}

func (ctrl *FileController) handleUploadFile(ctx *gin.Context) {
	uploadDto := &dtos.UploadFileDto{}
	if err := ctx.MustBindWith(uploadDto, binding.FormMultipart); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr)
		return
	}

	// fileHeader, err := ctx.FormFile("file")
	if url, err := ctrl.service.UploadFile(uploadDto.File); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr)
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(url))
	}
}
