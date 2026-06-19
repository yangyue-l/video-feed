package models

import "time"

// Comment 评论模型
type Comment struct {
	ID        int64     `json:"id,string" gorm:"primaryKey"`
	UserID    int64     `json:"user_id,string" gorm:"index;not null"`
	VideoID   int64     `json:"video_id,string" gorm:"index;not null"`
	ParentID  int64     `json:"parent_id,string" gorm:"default:0;index"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

// CommentResponse 评论响应
type CommentResponse struct {
	ID         int64        `json:"id,string"`
	User       UserResponse `json:"user"`
	Content    string       `json:"content"`
	CreateDate string       `json:"create_date"`
}

// TableName 表名
func (Comment) TableName() string {
	return "comments"
}
