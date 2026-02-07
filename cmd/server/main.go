package main

import (
	"fmt"
	"hmdp-go/internal/config"
	"hmdp-go/internal/controller"
	"hmdp-go/internal/repository"
	"hmdp-go/internal/router"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/mysql"
	"hmdp-go/pkg/redisutil"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release") // 设置 gin为发布模式
	
	config.InitConfig() // 初始化配置

	// 读取端口号
	port := config.AppConfig.Server.Port
	fmt.Printf("Server started on port %s\n", port)

	// 初始化数据库 - 返回数据库实例
	db := mysql.InitMySQL()

	// 初始化 Redis
	redisutil.InitRedis()

	// 初始化依赖
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	// 路由信息
	r := router.SetupRouter(ctrl)

	// 绑定端口,运行
	r.Run(":" + port)
}
