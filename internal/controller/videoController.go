package controller

import (
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	// TODO: 实现发布视频逻辑
	var p models.Video
	// 1. 获取用户信息
	userID := c.GetInt64("user_id")
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("PublishVideo with invalid param", zap.Error(err))
		utils.ParamError(c, Translate(err))
		return
	}
	user, err := dao.FindUserByID(userID)
	if err != nil {
		zap.L().Error("dao.FindUserByID() failed", zap.Error(err))
		utils.Error(c, 1)
		return
	}
	// 2. 获取上传的视频的信息
	_ = user // 暂时使用

	utils.Success(c, nil)
}

// GetVideoFeed 获取视频Feed流
func GetVideoFeed(c *gin.Context) {
	userID := c.GetInt64("user_id")

	videos, err := service.GetVideoFeed(userID, 0, 10)
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
	userID := c.GetInt64("user_id")

	videos, err := service.GetUserVideos(userID)
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
