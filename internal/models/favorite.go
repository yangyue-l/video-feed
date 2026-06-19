package models

import "time"

// Favorite 点赞模型
type Favorite struct {
	ID        int64     `json:"id,string" gorm:"primaryKey"`
	UserID    int64     `json:"user_id,string" gorm:"uniqueIndex:idx_user_video;not null"`
	VideoID   int64     `json:"video_id,string" gorm:"uniqueIndex:idx_user_video;not null"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Video     Video     `json:"video" gorm:"foreignKey:VideoID"`
}

// TableName 表名
func (Favorite) TableName() string {
	return "favorites"
}
