package result

import (
	"hmdp-go/internal/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 统一响应结构体
type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Success: true,
		Code:    constant.CodeSuccess,
		Msg:     constant.MsgSuccess,
		Data:    data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Success: false,
		Code:    code,
		Msg:     msg,
	})
}
