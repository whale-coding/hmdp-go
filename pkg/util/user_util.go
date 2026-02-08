package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserId 从Gin Context中获取用户ID
func GetUserId(c *gin.Context) (uint64, bool) {
	uid, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}
	userId, ok := uid.(uint64)
	return userId, ok
}

// StringToUint64 将字符串转换为 uint64
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
