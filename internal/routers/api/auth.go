package api

import (
	"github.com/gin-gonic/gin"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/service"
	"gotour/blog-service/pkg/app"
	"gotour/blog-service/pkg/errcode"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	ser := service.New(c.Request.Context())
	err := ser.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "service.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}