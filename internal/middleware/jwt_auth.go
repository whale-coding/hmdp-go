package middleware

import (
	"hmdp-go/internal/constant"
	"hmdp-go/pkg/jwt"
	"hmdp-go/pkg/redisutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JwtAuth 单Token+Redis鉴权中间件（含自动续期）
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. 接口白名单：登录、验证码等接口放行（路由那边分组了，现在不需要了）
		// whiteList := map[string]bool{
		// 	"/user/code":   true,
		// 	"/user/login":  true,
		// 	"/user/logout": true,
		// }
		// if whiteList[c.Request.URL.Path] {
		// 	c.Next()
		// 	return
		// }

		// 2. 从Header获取Token（规范：Bearer + 空格 + Token）
		authHeader := c.GetHeader("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头缺少 Authorization"})
			c.Abort()
			return
		}
		//parts := strings.SplitN(authHeader, " ", 2)
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token格式错误:Bearer + 空格 + Token"})
		//	c.Abort()
		//	return
		//}
		//tokenStr := parts[1] // 获取 Token部分
		tokenStr := authHeader

		// 3. 校验Token是否在黑名单（已退出的Token，防复用）
		blackKey := constant.LOGIN_BLACK_KEY + tokenStr
		if exists, _ := redisutil.Exists(blackKey); exists {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token已失效,请重新登录"})
			c.Abort()
			return
		}

		// 4. 解析JWT Token，获取用户ID
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token无效/已过期:" + err.Error()})
			c.Abort()
			return
		}
		userId := claims.UserId

		// 5. 校验Redis中是否存在该Token（防止JWT有效但已退出/过期）
		redisKey := constant.LOGIN_USER_KEY + tokenStr
		redisUserId, err := redisutil.Get(redisKey)
		if err != nil || redisUserId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token已过期/已退出"})
			c.Abort()
			return
		}

		// 6. 防Token伪造：校验Redis中的用户ID与JWT中的是否一致
		if redisUserId != strconv.FormatUint(userId, 10) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token无效,用户不匹配"})
			c.Abort()
			return
		}

		// 7. 自动续期！刷新Redis中Token的过期时间
		_ = redisutil.Expire(redisKey, constant.LOGIN_USER_TTL)

		// 8. 将用户ID存入Gin Context，业务接口直接获取
		c.Set("user_id", userId)

		// 9. 校验通过，放行请求
		c.Next()
	}
}
