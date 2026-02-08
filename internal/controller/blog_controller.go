package controller

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/result"
	"hmdp-go/pkg/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BlogController 博客控制器接口
type BlogController interface {
	QueryMyBlog(ctx *gin.Context)
}

// blogController 博客控制器实现
type blogController struct {
	svc *service.Service
}

// NewBlogController 创建博客控制器实例
func NewBlogController(svc *service.Service) BlogController {
	return &blogController{svc: svc}
}

// 查询当前登录用户的博客
func (c *blogController) QueryMyBlog(ctx *gin.Context) {
	// 1. 从Context获取用户ID（封装的工具函数）
	userId, ok := util.GetUserId(ctx)
	if !ok {
		result.Fail(ctx, constant.ErrCodeUnauthorized, "未获取到用户信息，请重新登录")
		return
	}
	log.Printf("获取当前登录用户的博客,登录用户ID: %d\n", userId)

	var page model.PaginationRequest
	// 2. 解析分页参数 current
	currentStr := ctx.Query("current")
	if current, err := strconv.Atoi(currentStr); err == nil && current > 0 {
		page.PageNo = current
	}
	// 设置默认值
	page.SetDefault()
	log.Printf("查询当前登录用户的博客,页码: %d,页大小: %d\n", page.PageNo, page.PageSize)

	// 3. 调用服务层查询博客列表
	res, err := c.svc.BlogService.QueryMyBlog(userId, &page)
	if err != nil {
		log.Printf("查询博客列表失败,用户ID: %d,错误: %v\n", userId, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, "查询博客列表失败")
		return
	}

	result.Success(ctx, res)
}
