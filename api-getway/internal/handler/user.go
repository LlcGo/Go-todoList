package handler

import (
	"api-getway/internal/logic"
	"api-getway/pkg/e"
	"api-getway/pkg/res"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ginCtx *gin.Context) {
	var userReq logic.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// gin.key 中获取服务实例
	userService := ginCtx.Keys["user"].(logic.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Status: uint(userResp.Code),
		Data:   userResp,
		Msg:    e.GetMsg(uint(userResp.Code)),
		Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UserLogin(ginCtx *gin.Context) {
	var userReq logic.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// gin.key 中获取服务实例
	userService := ginCtx.Keys["user"].(logic.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Status: uint(userResp.Code),
		Data:   userResp,
		Msg:    e.GetMsg(uint(userResp.Code)),
		Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}
