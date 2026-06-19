package service

import (
	"video_feed/internal/models"
)

// AddComment 添加评论
func AddComment(userID, videoID int64, content string) (*models.CommentResponse, error) {
	// TODO: 实现添加评论逻辑
	return nil, nil
}

// DeleteComment 删除评论
func DeleteComment(userID, commentID int64) error {
	// TODO: 实现删除评论逻辑
	return nil
}

// GetCommentList 获取视频评论列表
func GetCommentList(videoID int64) ([]models.CommentResponse, error) {
	// TODO: 实现获取评论列表逻辑
	return nil, nil
}
