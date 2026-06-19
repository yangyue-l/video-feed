package models

import "time"

// Video 视频模型
type Video struct {
	ID            int64     `json:"id,string" gorm:"primaryKey"`
	UserID        int64     `json:"user_id,string" gorm:"index;not null"`
	Title         string    `json:"title" gorm:"size:128;not null"`
	PlayURL       string    `json:"play_url" gorm:"size:256;not null"`
	CoverURL      string    `json:"cover_url" gorm:"size:256"`
	FavoriteCount int64     `json:"favorite_count" gorm:"default:0"`
	CommentCount  int64     `json:"comment_count" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
}

// VideoResponse 视频响应
type VideoResponse struct {
	ID            int64        `json:"id,string"`
	Author        UserResponse `json:"author"`
	PlayURL       string       `json:"play_url"`
	CoverURL      string       `json:"cover_url"`
	FavoriteCount int64        `json:"favorite_count"`
	CommentCount  int64        `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

// TableName 表名
func (Video) TableName() string {
	return "videos"
}
