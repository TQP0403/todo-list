package jwt

import (
	"TQP0403/todo-list/src/common"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const jwtKey = "user-id"

func (service *JwtService) JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request

		token := ""
		bearToken := req.Header.Get("Authorization")
		// //normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {
			token = strArr[1]
		}

		if data, err := service.JwtVerify(token); err != nil {
			cusErr := common.NewUnauthorizedError(err)
			ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr)
		} else {
			ctx.Set(jwtKey, data.UserId)
			ctx.Next()
		}
	}
}

func GetUserId(ctx *gin.Context) int {
	userId := ctx.MustGet(jwtKey).(int)
	if userId == 0 {
		cusErr := common.NewBadRequestError(errors.New("header userId not found"))
		ctx.AbortWithStatusJSON(cusErr.StatusCode, cusErr)
	}

	return userId
}
