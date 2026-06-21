package service

import (
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/utils"
)

// FollowUser 关注/取关用户
func FollowUser(userID, toUserID int64, actionType int) error {
	if userID == toUserID {
		return nil
	}

	if actionType == 1 {
		if dao.IsFollow(userID, toUserID) {
			return nil
		}
		relation := &models.Relation{
			ID:       utils.GenerateID(),
			UserID:   userID,
			ToUserID: toUserID,
		}
		return dao.CreateRelation(relation)
	}

	return dao.DeleteRelation(userID, toUserID)
}

// GetFollowList 获取关注列表
func GetFollowList(userID int64) ([]models.UserResponse, error) {
	users, err := dao.GetFollowList(userID)
	if err != nil {
		return nil, err
	}
	var res []models.UserResponse
	for _, u := range users {
		res = append(res, models.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
		})
	}
	return res, nil
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userID int64) ([]models.UserResponse, error) {
	users, err := dao.GetFollowerList(userID)
	if err != nil {
		return nil, err
	}
	var res []models.UserResponse
	for _, u := range users {
		res = append(res, models.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
		})
	}
	return res, nil
}
