package routes

import (
	"api-getway/internal/handler"
	"api-getway/middlerware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlerware.Cors())
	v1 := ginRouter.Group("api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)
	}
	return ginRouter
}
