package controller

import (
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// FavoriteVideo 点赞/取消点赞视频
func FavoriteVideo(c *gin.Context) {
	//1. 获取当前用户ID
	userID := c.GetInt64("user_id")

	//2. 绑定模型
	p := new(models.ParamFavorite)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.Error(c, 400)
			return
		}
		utils.ErrorWithMessage(c, 400, Translate(err))
		return
	}
	msg, err := service.FavoriteVideo(userID, p.VideoID, p.ActionType)
	if err != nil {
		zap.L().Error("service.FavouriteVideo failed",
			zap.Int64("userID", userID),
			zap.Int64("videoID", p.VideoID),
			zap.Error(err))
		utils.Error(c, 1)
		return
	}

	utils.Success(c, msg)
}

// GetFavoriteList 获取用户点赞列表
func GetFavoriteList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	videoList, err := service.GetFavoriteList(userID)
	if err != nil {
		zap.L().Error("service.GetFavoriteList() failed", zap.Error(err))
		utils.Error(c, 1)
		return
	}

	utils.Success(c, videoList)
}
