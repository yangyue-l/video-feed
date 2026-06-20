package service

import (
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/utils"
)

// AddComment 添加评论
func AddComment(userID, videoID, parentID int64, content string) (*models.CommentResponse, error) {
	// TODO: 实现添加评论逻辑
	//获取用户信息
	user, err := dao.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	//创建用户返回的对象
	userRes := &models.UserResponse{
		ID:       userID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	//创建评论对象
	comment := &models.Comment{
		ID:       utils.GenerateID(),
		UserID:   userID,
		VideoID:  videoID,
		ParentID: parentID,
		Content:  content,
	}
	if err := dao.CreateComment(comment); err != nil {
		return nil, err
	}

	commentRes := &models.CommentResponse{
		ID:         comment.ID,
		User:       *userRes,
		Content:    content,
		CreateDate: comment.CreatedAt.Format("2006-01-02"),
	}
	return commentRes, nil
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
