package controller

import (
	"strconv"
	"time"
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	var p models.ParamVideo
	// 1. 获取用户信息
	userID := c.GetInt64("user_id")

	// 2. 绑定表单参数（文件上传用 ShouldBind）
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("PublishVideo with invalid param", zap.Error(err))
		utils.ParamError(c, Translate(err))
		return
	}

	// 3. 获取视频文件
	videoFile, err := c.FormFile("video")
	if err != nil {
		zap.L().Error("get videoFile failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, "请上传视频文件")
		return
	}

	// 4. 获取封面图片
	coverFile, err := c.FormFile("coverFile")
	if err != nil {
		zap.L().Error("get coverFile failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, "请上传封面图片")
		return
	}

	// 5. 调用service发布视频
	if err := service.PublishVideo(userID, p.Title, videoFile, coverFile); err != nil {
		zap.L().Error("service.PublishVideo() failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, "发布成功")
}

// GetVideoFeed 获取视频Feed流
func GetVideoFeed(c *gin.Context) {
	userID := c.GetInt64("user_id")

	// 获取查询参数 latest_time
	var latestTime int64
	if t := c.Query("latest_time"); t != "" {
		latestTime, _ = strconv.ParseInt(t, 10, 64)
	} else {
		latestTime = time.Now().UnixMilli() // 默认当前时间戳(毫秒)
	}

	videos, err := service.GetVideoFeed(userID, latestTime, 10)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"video_list": videos,
	})
}

// GetUserVideos 获取用户视频列表
func GetUserVideos(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if userID == 0 {
		utils.Error(c, 400)
		return
	}

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
	userID := c.GetInt64("user_id")

	// 获取路径参数 video_id
	videoIDStr := c.Param("id")
	videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
	if err != nil {
		utils.ParamError(c, "视频ID格式错误")
		return
	}

	video, err := service.GetVideoDetail(userID, videoID)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, video)
}
