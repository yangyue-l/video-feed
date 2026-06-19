package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"
)

// CreateComment 创建评论
func CreateComment(comment *models.Comment) error {
	return database.DB.Create(comment).Error
}

// DeleteComment 删除评论
func DeleteComment(id int64) error {
	return database.DB.Delete(&models.Comment{}, id).Error
}

// FindCommentByID 根据ID查找评论
func FindCommentByID(id int64) (*models.Comment, error) {
	var comment models.Comment
	err := database.DB.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentList 获取视频评论列表
func GetCommentList(videoID int64) ([]models.Comment, error) {
	var comments []models.Comment
	err := database.DB.Where("video_id = ?", videoID).Order("created_at DESC").Find(&comments).Error
	return comments, err
}
