package service

import (
	"errors"
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// Register 用户注册
func Register(username, password string) error {
	// 1. 检查用户名是否已存在
	if err := dao.CheckUserExist(username); err != nil {
		return err
	}
	// 2. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("服务器繁忙")
	}
	// 3. 生成分布式ID
	userID := utils.GenerateID()
	// 4. 创建用户
	user := &models.User{
		ID:       userID,
		Username: username,
		Password: string(hashedPassword),
		Nickname: username,
	}
	return dao.InsertUser(user)
}

// Login 用户登录
func Login(username, password string) (*models.TokenResponse, error) {
	// 1. 查找用户
	user, err := dao.FindUserByUsername(username)
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
func GetUserInfo(userID int64) (*models.User, error) {
	return dao.FindUserByID(userID)
}
