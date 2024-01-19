package auth

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"TQP0403/todo-list/src/modules/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service    IAuthService
	jwtService jwt.IJwtService
}

func NewController(service *AuthService, jwtService *jwt.JwtService) *AuthController {
	return &AuthController{service: service, jwtService: jwtService}
}

func (ctrl *AuthController) Register(router *gin.Engine) {
	group := router.Group("/api/auth")
	{
		group.POST("/register", ctrl.handleRegister)
		group.POST("/login", ctrl.handleLogin)
		group.POST("/refresh-token", ctrl.handleRefreshToken)
		group.GET("/profile", jwt.JwtMiddleware(ctrl.jwtService), ctrl.handleGetProfile)
	}
}

func (ctrl *AuthController) handleRefreshToken(ctx *gin.Context) {
	var res *LoginResponse
	var reqData dtos.RefreshTokenDto
	var err error

	if err := ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	if res, err = ctrl.service.RefreshToken(reqData.Token); err != nil {
		cusErr := common.NewUnauthorizedError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	// return JWT
	ctx.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (ctrl *AuthController) handleLogin(ctx *gin.Context) {
	var res *LoginResponse
	var reqData dtos.LoginDto
	var err error

	if err = ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	if res, err = ctrl.service.Login(&reqData); err != nil {
		cusErr := common.NewUnauthorizedError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	// return JWT
	ctx.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (ctrl *AuthController) handleRegister(ctx *gin.Context) {
	var res *LoginResponse
	var reqData dtos.RegisterDto
	var err error

	if err = ctx.ShouldBind(&reqData); err != nil {
		cusErr := common.NewBadRequestError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	if res, err = ctrl.service.Register(&reqData); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (ctrl *AuthController) handleGetProfile(ctx *gin.Context) {
	userId := jwt.GetUserId(ctx)

	if res, err := ctrl.service.GetProfile(userId); err != nil {
		cusErr := common.NewInternalServerError(err)
		ctx.JSON(cusErr.StatusCode, common.NewErrorResponse(*cusErr))
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(res))
	}
}
