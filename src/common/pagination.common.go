package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Total    int64 `json:"total" form:"-"`
	Page     int   `json:"page" form:"page"`
	PageSize int   `json:"pageSize" form:"page-size"`
}

type PaginationResponse struct {
	Metadata Pagination  `json:"metadata"`
	Rows     interface{} `json:"rows"`
}

func BindPagination(ctx *gin.Context) Pagination {
	query := Pagination{Page: 1, PageSize: 10}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   "BadRequest",
		})
	}

	if query.Page < 1 {
		query.Page = 1
	}

	if query.PageSize < 1 {
		query.PageSize = 10
	}

	return query
}
