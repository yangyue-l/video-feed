package dao

import (
	"errors"
	"video_feed/internal/database"
	"video_feed/internal/models"
)

// CheckUserExist 检查用户名是否已存在
func CheckUserExist(username string) error {
	var count int64
	database.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("用户名已经存在")
	}
	return nil
}

// InsertUser 插入用户
func InsertUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// FindByUsername 根据用户名查找用户
func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找用户
func FindUserByID(id int64) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户
func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}
