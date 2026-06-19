package service

import (
	"video_feed/internal/models"
)

// FavoriteVideo 点赞/取消点赞视频
func FavoriteVideo(userID, videoID int64, actionType int) error {
	// TODO: 实现点赞/取消点赞逻辑
	return nil
}

// GetFavoriteList 获取用户点赞列表
func GetFavoriteList(userID int64) ([]models.VideoResponse, error) {
	// TODO: 实现获取点赞列表逻辑
	return nil, nil
}
