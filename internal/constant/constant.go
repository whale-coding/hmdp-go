package constant

import "time"

// 错误码
const (
	CodeSuccess           = 200
	ErrCodeInvalidParam   = 400
	ErrCodeUserExist      = 401
	ErrCodeUserNotFound   = 402
	ErrCodePasswordError  = 403
	ErrCodeTokenInvalid   = 404
	ErrCodeServerInternal = 500
)

// 错误信息
const (
	MsgSuccess        = "success"
	MsgInvalidParam   = "参数错误"
	MsgUserExist      = "用户已存在"
	MsgUserNotFound   = "用户不存在"
	MsgPasswordError  = "密码错误"
	MsgTokenInvalid   = "token无效"
	MsgServerInternal = "服务器内部错误"
)

// Redis Key 前缀
const (
	LOGIN_CODE_KEY = "login:code:"  // 登录验证码的Redis键前缀
	LOGIN_USER_KEY = "login:token:" // 登录用户的Redis键前缀
	CACHE_SHOP_KEY = "cache:shop:"  // 商铺缓存的Redis键前缀
)

// Redis TTL 过期时间
const (
	LOGIN_CODE_TTL = 2 * time.Minute  // 登录验证码过期时间
	LOGIN_USER_TTL = 30 * time.Minute // 登录用户过期时间
	CACHE_SHOP_TTL = 30 * time.Minute // 商铺缓存过期时间
)
