package models

import "time"

// Comment 评论模型
type Comment struct {
	// TODO: 定义评论模型字段
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	VideoID   uint      `json:"video_id" gorm:"index;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
}

// CommentResponse 评论响应
type CommentResponse struct {
	// TODO: 定义评论响应字段
	ID         uint         `json:"id"`
	User       UserResponse `json:"user"`
	Content    string       `json:"content"`
	CreateDate string       `json:"create_date"`
}

// TableName 表名
func (Comment) TableName() string {
	return "comments"
}
