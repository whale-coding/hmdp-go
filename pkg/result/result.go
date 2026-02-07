package result

import (
	"hmdp-go/internal/constant"

	"github.com/gin-gonic/gin"
)

// Result 统一响应结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Result{
		Code: constant.ErrCodeSuccess,
		Msg:  constant.MsgSuccess,
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, Result{
		Code: code,
		Msg:  msg,
	})
}
