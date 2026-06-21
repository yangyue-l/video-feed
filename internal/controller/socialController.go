package controller

import (
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FollowUser 关注/取关用户
func FollowUser(c *gin.Context) {
	userID := c.GetInt64("user_id")

	p := new(models.ParamUserAction)
	if err := c.ShouldBindJSON(p); err != nil {
		utils.Error(c, 400)
		return
	}

	if err := service.FollowUser(userID, p.ToUserID, p.ActionType); err != nil {
		zap.L().Error("service.FollowUser() failed", zap.Error(err))
		utils.Error(c, 500)
		return
	}

	utils.Success(c, nil)
}

// GetFollowList 获取关注列表
func GetFollowList(c *gin.Context) {
	userID := c.GetInt64("user_id")

	users, err := service.GetFollowList(userID)
	if err != nil {
		zap.L().Error("service.GetFollowList() failed", zap.Error(err))
		utils.Error(c, 500)
		return
	}

	utils.Success(c, users)
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c *gin.Context) {
	userID := c.GetInt64("user_id")

	users, err := service.GetFollowerList(userID)
	if err != nil {
		zap.L().Error("service.GetFollowerList() failed", zap.Error(err))
		utils.Error(c, 500)
		return
	}

	utils.Success(c, users)
}
