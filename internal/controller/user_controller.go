package controller

import (
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/service"
	"hmdp-go/pkg/result"
	"hmdp-go/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器接口
type UserController interface {
	SendCode(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Me(ctx *gin.Context)
	Info(ctx *gin.Context)
}

// userController 用户控制器实现
type userController struct {
	svc *service.Service
}

// NewUserController 创建用户控制器实例
func NewUserController(svc *service.Service) UserController {
	return &userController{svc: svc}
}

// SendCode 发送手机验证码
func (c *userController) SendCode(ctx *gin.Context) {
	// 从查询参数中获取手机号
	phone := ctx.Query("phone")
	if phone == "" {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "手机号不能为空")
		return
	}

	// 调用服务层发送验证码
	err := c.svc.UserService.SendCode(phone)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, err.Error())
		return
	}

	result.Success(ctx, nil)
}

// Login 登录
func (c *userController) Login(ctx *gin.Context) {
	// 获取请求参数
	var req model.LoginFormDTO

	// 参数绑定与验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "请求参数错误")
		return
	}
	log.Printf("登录请求参数: %+v\n", req)

	// 调用服务层的登录逻辑
	token, err := c.svc.UserService.Login(&req)
	if err != nil {
		log.Printf("登录失败: %v\n", err)
		result.Fail(ctx, constant.ErrCodeInvalidParam, err.Error())
		return
	}
	log.Printf("登录成功,生成的Token: %s\n", token)

	result.Success(ctx, token)
}

// Logout 登出
func (c *userController) Logout(ctx *gin.Context) {
	// 1. 从Header获取Token
	token := ctx.GetHeader("authorization")
	if token == "" {
		result.Fail(ctx, constant.ErrCodeUnauthorized, "请求头缺少 Authorization")
		return
	}

	// 2. 调用服务层登出
	_ = c.svc.UserService.Logout(token)

	// 3. 无论Token是否有效，都返回成功（前端清除本地状态即可）
	result.Success(ctx, nil)
}

// Me 获取当前登录用户信息
func (c *userController) Me(ctx *gin.Context) {
	// 1. 从Context获取用户ID（封装的工具函数）
	userId, ok := util.GetUserId(ctx)
	if !ok {
		result.Fail(ctx, constant.ErrCodeUnauthorized, "未获取到用户信息，请重新登录")
		return
	}
	log.Printf("获取当前登录用户信息,登录用户ID: %d\n", userId)
	// 调用服务层的逻辑
	res, err := c.svc.UserService.GetLoginUser(userId)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeServerInternal, "获取用户信息失败")
		return
	}

	result.Success(ctx, res)
}

// Info 获取用户信息
func (c *userController) Info(ctx *gin.Context) {
	// 1. 从路径参数获取用户ID
	idStr := ctx.Param("id")
	id, err := util.StringToUint64(idStr)
	if err != nil {
		result.Fail(ctx, constant.ErrCodeInvalidParam, "用户ID参数错误")
		return
	}
	log.Printf("获取用户信息,用户ID: %d\n", id)

	// 2. 调用服务层查询用户信息
	res, err := c.svc.UserInfoService.GetUserInfo(id)
	if err != nil {
		log.Printf("获取用户信息失败,用户ID: %d,错误信息: %v\n", id, err)
		result.Fail(ctx, constant.ErrCodeServerInternal, err.Error())
		return
	}
	log.Printf("获取用户信息成功,用户ID: %d,用户信息: %+v\n", id, res)

	result.Success(ctx, res)
}
