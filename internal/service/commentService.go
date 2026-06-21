package service

import (
	"fmt"
	"video_feed/internal/dao"
	"video_feed/internal/database"
	"video_feed/internal/models"
	"video_feed/internal/utils"

	"gorm.io/gorm"
)

// AddComment 添加评论
func AddComment(userID, videoID, parentID int64, content string) (*models.CommentResponse, error) {
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

	if err := dao.UpdateCommentCountTx(database.DB, videoID, 1); err != nil {
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
	return database.DB.Transaction(func(tx *gorm.DB) error {
		comment, err := dao.FindCommentByIDTx(tx, commentID)
		if err != nil {
			return err
		}

		if comment.UserID != userID {
			video, err := dao.FindVideoByIDTx(tx, comment.VideoID)
			if err != nil {
				return err
			}
			if video.UserID != userID {
				return fmt.Errorf("无权删除该评论")
			}
		}

		if err := dao.DeleteCommentTx(tx, commentID); err != nil {
			return err
		}

		return dao.UpdateCommentCountTx(tx, comment.VideoID, -1)
	})
}

// GetCommentList 获取视频评论列表
func GetCommentList(videoID int64) ([]models.CommentResponse, error) {
	comments, err := dao.GetCommentList(videoID)
	if err != nil {
		return nil, err
	}
	var commentResS []models.CommentResponse
	for _, comment := range comments {
		commentRes := &models.CommentResponse{
			ID: comment.ID,
			User: models.UserResponse{
				ID:       comment.User.ID,
				Username: comment.User.Username,
				Nickname: comment.User.Nickname,
				Avatar:   comment.User.Avatar,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02"),
		}
		commentResS = append(commentResS, *commentRes)
	}
	return commentResS, nil
}
