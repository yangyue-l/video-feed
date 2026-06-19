package controller

import (
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
)

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	// TODO: 实现发布视频逻辑
	// 1. 获取视频文件和封面文件
	// 2. 验证文件类型和大小
	// 3. 保存文件到服务器
	// 4. 创建视频记录
	// 5. 返回响应

	utils.Success(c, nil)
}

// GetVideoFeed 获取视频Feed流
func GetVideoFeed(c *gin.Context) {
	userID := c.GetUint("user_id")

	videoService := service.NewVideoService()
	videos, err := videoService.GetVideoFeed(userID, 0, 10)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"videos": videos,
	})
}

// GetUserVideos 获取用户视频列表
func GetUserVideos(c *gin.Context) {
	userID := c.GetUint("user_id")

	videoService := service.NewVideoService()
	videos, err := videoService.GetUserVideos(userID)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"videos": videos,
	})
}

// GetVideoDetail 获取视频详情
func GetVideoDetail(c *gin.Context) {
	// TODO: 实现获取视频详情逻辑
	utils.Success(c, nil)
}
