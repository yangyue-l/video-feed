package service

import (
	"video_feed/internal/models"
)

// FollowUser 关注/取关用户
func FollowUser(userID, toUserID int64, actionType int) error {
	// TODO: 实现关注/取关逻辑
	return nil
}

// GetFollowList 获取关注列表
func GetFollowList(userID int64) ([]models.UserResponse, error) {
	// TODO: 实现获取关注列表逻辑
	return nil, nil
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userID int64) ([]models.UserResponse, error) {
	// TODO: 实现获取粉丝列表逻辑
	return nil, nil
}
