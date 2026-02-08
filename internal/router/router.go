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
	api := r.Group("")

	// 公开路由（无需登录）
	user := api.Group("/user")
	{
		user.POST("/code", ctrl.UserController.SendCode)
		user.POST("/login", ctrl.UserController.Login)
		user.POST("/logout", ctrl.UserController.Logout)
	}

	// 需要登录的路由（使用JWT鉴权中间件）
	authApi := api.Group("")
	authApi.Use(middleware.JwtAuth())
	{
		authUser := authApi.Group("/user")
		{
			authUser.GET("/me", ctrl.UserController.Me)
			authUser.GET("/info/:id", ctrl.UserController.Info)
		}

		authBlog := authApi.Group("/blog")
		{
			authBlog.GET("/of/me", ctrl.BlogController.QueryMyBlog)
			authBlog.GET("/hot", ctrl.BlogController.QueryHotBlog)
			authBlog.GET("/:id", ctrl.BlogController.QueryBlogById)
		}

		authShopType := authApi.Group("/shop-type")
		{
			authShopType.GET("/list", ctrl.ShopTypeController.GetShopTypeList)
		}
	}

	// 返回 gin引擎
	return r
}
