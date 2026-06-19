package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"
)

// CreateFavorite 创建点赞
func CreateFavorite(favorite *models.Favorite) error {
	return database.DB.Create(favorite).Error
}

// DeleteFavorite 删除点赞
func DeleteFavorite(userID, videoID int64) error {
	return database.DB.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&models.Favorite{}).Error
}

// IsFavorite 检查是否已点赞
func IsFavorite(userID, videoID int64) bool {
	var count int64
	database.DB.Model(&models.Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Count(&count)
	return count > 0
}

// GetFavoriteList 获取用户点赞列表
func GetFavoriteList(userID int64) ([]models.Video, error) {
	var videos []models.Video
	err := database.DB.Joins("JOIN favorites ON favorites.video_id = videos.id").
		Where("favorites.user_id = ?", userID).
		Order("favorites.created_at DESC").
		Find(&videos).Error
	return videos, err
}
