package service

import (
	"errors"
	"fmt"
	"hmdp-go/internal/constant"
	"hmdp-go/internal/model"
	"hmdp-go/internal/repository"
	"hmdp-go/pkg/jwt"
	"hmdp-go/pkg/redisutil"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	SendCode(phone string) error
	Login(req *model.LoginFormDTO) (string, error)
	Logout(token string) error
	GetLoginUser(userId uint64) (*model.UserDTO, error)
}

// userService 用户服务实现
type userService struct {
	repo *repository.Repository
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repository.Repository) UserService {
	return &userService{repo: repo}
}

// SendCode 发送验证码
func (s *userService) SendCode(phone string) error {
	// 1. 校验手机号格式
	if !isValidPhone(phone) {
		return errors.New("手机号格式不正确")
	}

	// 2. 生成6位随机验证码
	code := generateCode()

	// 3. 保存验证码到Redis
	loginCodeKey := constant.LOGIN_CODE_KEY + phone
	err := redisutil.Set(loginCodeKey, code, constant.LOGIN_CODE_TTL)
	if err != nil {
		return errors.New("验证码保存失败")
	}

	// 4. 模拟发送验证码（实际项目中应调用短信服务）
	fmt.Printf("发送验证码到 %s: %s\n", phone, code)

	return nil
}

// Login 用户登录
func (s *userService) Login(req *model.LoginFormDTO) (string, error) {
	// 1. 校验手机号格式
	if !isValidPhone(req.Phone) {
		return "", errors.New("手机号格式不正确")
	}

	// 2. 从Redis获取验证码并校验
	loginCodeKey := constant.LOGIN_CODE_KEY + req.Phone
	cacheCode, err := redisutil.Get(loginCodeKey)
	if err != nil {
		return "", errors.New("验证码已过期，请重新获取")
	}
	if cacheCode != req.Code {
		return "", errors.New("验证码错误")
	}

	// 3. 根据手机号查询用户，不存在则创建新用户
	user, err := s.repo.UserRepo.FindByPhone(req.Phone)
	if err != nil {
		log.Printf("查询用户结果: err=%v, 是否为ErrRecordNotFound=%v\n", err, errors.Is(err, gorm.ErrRecordNotFound))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在，创建新用户
			user = &model.User{
				Phone:    req.Phone,
				NickName: "user_" + generateCode(), // 随机生成昵称
				Icon:     "",
			}
			if err := s.repo.UserRepo.Create(user); err != nil {
				log.Printf("创建用户失败: %v\n", err)
				return "", errors.New("创建用户失败")
			}
			log.Printf("新用户注册成功: %s (ID: %d)\n", user.Phone, user.ID)
		} else {
			log.Printf("查询用户失败(非NotFound): %v\n", err)
			return "", errors.New("查询用户失败")
		}
	} else {
		log.Printf("用户已存在: %s (ID: %d)\n", user.Phone, user.ID)
	}

	// 4. 生成jwt token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("生成Token失败")
	}

	// 5. 将用户ID保存到Redis（用于防伪Token校验）
	tokenKey := constant.LOGIN_USER_KEY + token
	err = redisutil.Set(tokenKey, strconv.FormatUint(user.ID, 10), constant.LOGIN_USER_TTL)
	if err != nil {
		return "", errors.New("保存登录状态失败")
	}

	// 6. 删除已使用的验证码
	_ = redisutil.Del(loginCodeKey)

	// 7. 返回token
	return token, nil
}

// Logout 用户登出
func (s *userService) Logout(token string) error {
	if token == "" {
		return nil
	}

	// 1. 从Redis删除Token记录
	tokenKey := constant.LOGIN_USER_KEY + token
	_ = redisutil.Del(tokenKey)

	// 2. 将Token加入黑名单（防止JWT有效期内继续使用）
	// 每个 token 独立存储，过期时间设为 JWT 最大有效期
	blackKey := constant.LOGIN_BLACK_KEY + token
	_ = redisutil.Set(blackKey, "1", constant.LOGIN_USER_TTL)

	log.Println("用户退出登录，Token加入黑名单:", token)

	return nil
}

// GetLoginUser 获取当前登录用户信息
func (s *userService) GetLoginUser(userId uint64) (*model.UserDTO, error) {
	// 1. 根据用户ID查询用户
	user, err := s.repo.UserRepo.FindByID(userId)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 2. 转换为DTO返回（隐藏敏感信息）
	userDTO := &model.UserDTO{
		ID:       user.ID,
		NickName: user.NickName,
		Icon:     user.Icon,
	}

	return userDTO, nil
}

// isValidPhone 校验手机号格式
func isValidPhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

// generateCode 生成6位随机验证码
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
