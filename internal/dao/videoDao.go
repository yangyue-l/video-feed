package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// CreateVideo 创建视频
func CreateVideo(video *models.Video) error {
	return database.DB.Create(video).Error
}

// FindVideoByID 根据ID查找视频
func FindVideoByID(id int64) (*models.Video, error) {
	var video models.Video
	err := database.DB.First(&video, id).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

// GetVideoFeed 获取视频Feed流
func GetVideoFeed(latestTime int64, limit int) ([]models.Video, error) {
	var videos []models.Video
	query := database.DB.Order("created_at DESC")
	if latestTime > 0 {
		query = query.Where("created_at < ?", latestTime)
	}
	err := query.Limit(limit).Find(&videos).Error
	return videos, err
}

// GetUserVideos 获取用户发布的视频
func GetUserVideos(userID int64) ([]models.Video, error) {
	var videos []models.Video
	err := database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&videos).Error
	return videos, err
}

// UpdateFavoriteCount 更新点赞数
func UpdateFavoriteCount(videoID int64, delta int64) error {
	return database.DB.Model(&models.Video{}).Where("id = ?", videoID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", delta)).Error
}

// UpdateCommentCount 更新评论数
func UpdateCommentCount(videoID int64, delta int64) error {
	return database.DB.Model(&models.Video{}).Where("id = ?", videoID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", delta)).Error
}
