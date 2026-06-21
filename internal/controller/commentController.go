package controller

import (
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {
	// TODO: 实现添加评论逻辑
	userID := c.GetInt64("user_id")

	p := new(models.ParamCommentAction)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.Error(c, 400)
			return
		}
		utils.ErrorWithMessage(c, 400, Translate(err))
		return
	}
	comment, err := service.AddComment(userID, p.VideoID, p.ParentID, p.Content)
	if err != nil {
		zap.L().Error("service.AddComment() failed", zap.Error(err))
		utils.Error(c, 500)
		return
	}

	utils.Success(c, comment)
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	userID := c.GetInt64("user_id")

	p := new(models.ParamDeleteComment)
	if err := c.ShouldBindUri(p); err != nil {
		utils.Error(c, 400)
		return
	}

	if err := service.DeleteComment(userID, p.CommentID); err != nil {
		zap.L().Error("service.DeleteComment() failed", zap.Error(err))
		utils.ErrorWithMessage(c, 403, err.Error())
		return
	}

	utils.Success(c, nil)
}

// GetCommentList 获取视频评论列表
func GetCommentList(c *gin.Context) {
	// TODO: 实现获取评论列表逻辑

	utils.Success(c, nil)
}
