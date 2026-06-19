package controller

import (
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
)

// FollowUser 关注/取关用户
func FollowUser(c *gin.Context) {
	// TODO: 实现关注/取关逻辑
	utils.Success(c, nil)
}

// GetFollowList 获取关注列表
func GetFollowList(c *gin.Context) {
	userID := c.GetUint("user_id")
	_ = userID

	// TODO: 调用service层获取关注列表
	utils.Success(c, nil)
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c *gin.Context) {
	userID := c.GetUint("user_id")
	_ = userID

	// TODO: 调用service层获取粉丝列表
	utils.Success(c, nil)
}
