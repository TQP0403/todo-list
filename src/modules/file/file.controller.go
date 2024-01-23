package file

import (
	"TQP0403/todo-list/src/common"
	"net/http"

	"github.com/gin-gonic/gin"
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
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
	}

	url, err := ctrl.service.UploadFile(fileHeader)
	if err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(cusErr))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
