package service

import (
	"video_feed/internal/dao"
	"video_feed/internal/models"
)

// UserService 用户服务层
type UserService struct {
	userDao *dao.UserDao
}

// NewUserService 创建用户服务
func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

// Register 用户注册
func (s *UserService) Register(username, password string) (*models.TokenResponse, error) {
	// TODO: 实现用户注册逻辑
	// 1. 检查用户名是否已存在
	// 2. 密码加密
	// 3. 创建用户
	// 4. 生成Token
	// 5. 返回Token

	return nil, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*models.TokenResponse, error) {
	// TODO: 实现用户登录逻辑
	// 1. 查找用户
	// 2. 验证密码
	// 3. 生成Token
	// 4. 返回Token

	return nil, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID uint) (*models.UserResponse, error) {
	// TODO: 实现获取用户信息逻辑
	// 1. 查找用户
	// 2. 获取关注数和粉丝数
	// 3. 组装响应数据

	return nil, nil
}
