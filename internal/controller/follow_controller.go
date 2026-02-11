package controller

import (
	"hmdp-go/internal/service"
)

// FollowController 关注控制器接口
type FollowController interface {
}

// followController 关注控制器实现
type followController struct {
	svc *service.Service
}

// NewFollowController 创建关注控制器实例
func NewFollowController(svc *service.Service) FollowController {
	return &followController{svc: svc}
}

// 关注和取关 @PutMapping("/{id}/{isFollow}")
// func (c *followController) Follow() {
// 	// 1. 获取用户ID
// 	userId, ok := util.GetUserId(ctx)
// 	if !ok {
// 		result.Fail(ctx, constant.ErrCodeUnauthorized, "未获取到用户信息，请重新登录")
// 		return
// 	}
// 	log.Printf("用户%d开始关注用户%d\n", userId, id)
// 	// 2. 获取关注状态
// 	isFollow := ctx.Param("isFollow") == "true"
// 	// 3. 执行关注或取关操作
// 	if err := c.svc.FollowService.FollowOrUnfollow(ctx, userId, id, isFollow); err != nil {
// 		result.Fail(ctx, constant.ErrCodeInternalServerError, "操作失败")
// 		return
// 	}
// 	action := "关注"
// 	if !isFollow {
// 		action = "取关"
// 	}
// 	log.Printf("用户%d成功%s用户%d\n", userId, action, id)
// 	result.Ok(ctx)
// }

// 当前登录用户是否关注了博主
// func (c *followController) IsFollow() {
// 	// 1. 获取用户ID
// 	userId, ok := util.GetUserId(ctx)
// 	if !ok {
// 		result.Fail(ctx, constant.ErrCodeUnauthorized, "未获取到用户信息，请重新登录")
// 		return
// 	}
// 	// 2. 获取博主ID
// 	id := ctx.Param("id")
// 	// 3. 查询关注状态
// 	isFollow, err := c.svc.FollowService.IsFollow(ctx, userId, id)
// 	if err != nil {
// 		result.Fail(ctx, constant.ErrCodeInternalServerError, "查询失败")
// 		return
// 	}
// 	result.OkWithData(ctx, isFollow)
// }

// 共同关注 @GetMapping("/common/{id}")
// func (c *followController) CommonFollow() {
// 	// 1. 获取用户ID
// 	userId, ok := util.GetUserId(ctx)
// 	if !ok {
// 		result.Fail(ctx, constant.ErrCodeUnauthorized, "未获取到用户信息，请重新登录")
// 		return
// 	}
// 	// 2. 获取博主ID
// 	id := ctx.Param("id")
// 	// 3. 查询共同关注列表
// 	commonFollows, err := c.svc.FollowService.CommonFollow(ctx, userId, id)
// 	if err != nil {
// 		result.Fail(ctx, constant.ErrCodeInternalServerError, "查询失败")
// 		return
// 	}
// 	result.OkWithData(ctx, commonFollows)
// }
