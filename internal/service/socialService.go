package service

import (
	"video_feed/internal/dao"
	"video_feed/internal/models"
)

// SocialService 社交服务层
type SocialService struct {
	relationDao *dao.RelationDao
	userDao     *dao.UserDao
}

// NewSocialService 创建社交服务
func NewSocialService() *SocialService {
	return &SocialService{
		relationDao: dao.NewRelationDao(),
		userDao:     dao.NewUserDao(),
	}
}

// FollowUser 关注/取关用户
func (s *SocialService) FollowUser(userID, toUserID uint, actionType int) error {
	// TODO: 实现关注/取关逻辑
	// 1. 不能关注自己
	// 2. 检查目标用户是否存在
	// 3. 关注时检查是否已关注
	// 4. 取关时检查是否已关注

	return nil
}

// GetFollowList 获取关注列表
func (s *SocialService) GetFollowList(userID uint) ([]models.UserResponse, error) {
	// TODO: 实现获取关注列表逻辑

	return nil, nil
}

// GetFollowerList 获取粉丝列表
func (s *SocialService) GetFollowerList(userID uint) ([]models.UserResponse, error) {
	// TODO: 实现获取粉丝列表逻辑

	return nil, nil
}
