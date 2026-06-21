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

// FindUserByVideoID 根据发布视频的用户ID找到改用户
func FindUserByVideoID(videoID int64) (u *models.User, err error) {
	u = new(models.User)
	err = database.DB.Joins("JOIN videos on videos.user_id = users.id").
		Where("videos.id = ?", videoID).
		First(u).Error
	return
}

// FindUserByCommentID 根据评论ID找到发布评论的用户
func FindUserByCommentID(commentID int64) (u *models.User, err error) {
	u = new(models.User)
	err = database.DB.Joins("JOIN comments on comments.user_id = users.id").
		Where("comments.id = ?", commentID).
		First(u).Error
	return
}
