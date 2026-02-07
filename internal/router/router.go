package router

import (
	"hmdp-go/internal/controller"
	"hmdp-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(ctrl *controller.Controller) *gin.Engine {
	// 初始化 gin引擎
	r := gin.Default()

	// 跨域配置需要放在路由前面，否则不生效！
	r.Use(middleware.Cors()) // 跨域中间件

	// API 版本
	//api := r.Group("/api")

	// 认证相关路由 (公开)
	//user := api.Group("/user")
	{
		//user.POST("/code", ctrl.UserController.SendCode)
		//user.POST("/login", ctrl.UserController.Login)
		//user.GET("/logout", ctrl.UserController.Logout)
	}

	// 返回 gin引擎
	return r
}
