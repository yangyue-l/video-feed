package service

import (
	"errors"
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/utils"

	"golang.org/x/crypto/bcrypt"
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
func (s *UserService) Register(username, password string) error {
	// 1. 检查用户名是否已存在
	if err := s.userDao.FindUserByName(username); err != nil {
		return errors.New("用户名已经存在")
	}
	// 2. 密码加密
	oPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("服务器繁忙")
	}
	// 3.生成分布式id
	userID := utils.GenerateID()
	// 3. 创建用户
	user := &models.User{
		ID:       userID,
		Username: username,
		Password: string(oPassword),
		Nickname: username,
	}
	return s.userDao.Save(user)
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*models.TokenResponse, error) {
	// 1. 查找用户
	user, err := s.userDao.FindByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 2. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	// 3. 生成token
	token, err := utils.GenerateToken(user.ID, username)
	if err != nil {
		return nil, errors.New("服务器繁忙")
	}
	return &models.TokenResponse{
		Token: token,
	}, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID int64) (*models.UserResponse, error) {
	// TODO: 实现获取用户信息逻辑
	// 1. 查找用户

	// 2. 获取关注数和粉丝数
	// 3. 组装响应数据

	return nil, nil
}
