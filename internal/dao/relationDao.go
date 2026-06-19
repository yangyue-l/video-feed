package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"
)

// CreateRelation 创建关注关系
func CreateRelation(relation *models.Relation) error {
	return database.DB.Create(relation).Error
}

// DeleteRelation 删除关注关系
func DeleteRelation(userID, toUserID int64) error {
	return database.DB.Where("user_id = ? AND to_user_id = ?", userID, toUserID).Delete(&models.Relation{}).Error
}

// IsFollow 检查是否已关注
func IsFollow(userID, toUserID int64) bool {
	var count int64
	database.DB.Model(&models.Relation{}).Where("user_id = ? AND to_user_id = ?", userID, toUserID).Count(&count)
	return count > 0
}

// GetFollowList 获取关注列表
func GetFollowList(userID int64) ([]models.User, error) {
	var users []models.User
	err := database.DB.Joins("JOIN relations ON relations.to_user_id = users.id").
		Where("relations.user_id = ?", userID).
		Order("relations.created_at DESC").
		Find(&users).Error
	return users, err
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userID int64) ([]models.User, error) {
	var users []models.User
	err := database.DB.Joins("JOIN relations ON relations.user_id = users.id").
		Where("relations.to_user_id = ?", userID).
		Order("relations.created_at DESC").
		Find(&users).Error
	return users, err
}
