package models

import "time"

// Favorite 点赞模型
type Favorite struct {
	// TODO: 定义点赞模型字段
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	VideoID   uint      `json:"video_id" gorm:"index;not null"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 表名
func (Favorite) TableName() string {
	return "favorites"
}
