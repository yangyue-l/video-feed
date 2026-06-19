package controller

import (
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
)

// FavoriteVideo 点赞/取消点赞视频
func FavoriteVideo(c *gin.Context) {
	// TODO: 实现点赞/取消点赞逻辑
	utils.Success(c, nil)
}

// GetFavoriteList 获取用户点赞列表
func GetFavoriteList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	_ = userID

	// TODO: 调用service层获取点赞列表
	utils.Success(c, nil)
}
