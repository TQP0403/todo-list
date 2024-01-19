package jwt

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware(jwtService IJwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request

		token := ""
		bearToken := req.Header.Get("Authorization")
		// //normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {
			token = strArr[1]
		}

		if data, err := jwtService.JwtVerify(token); err != nil {
			cusErr := common.NewUnauthorizedError(err)
			ctx.AbortWithStatusJSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		} else {
			// jsonStr, _ := json.Marshal(data)
			// ctx.Header("user", string(jsonStr))
			ctx.Request.Header.Add("user-id", fmt.Sprint(data.UserId))
			ctx.Next()
		}
	}
}

func GetUserId(ctx *gin.Context) int {
	userIdStr := ctx.Request.Header.Get("user-id")
	userId := helper.ParseInt(userIdStr)
	if userId == 0 {
		cusErr := common.NewBadRequestError(errors.New("header userId not found"))
		ctx.AbortWithStatusJSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
	}

	return userId
}
