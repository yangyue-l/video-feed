package models

import "time"

// Relation 关注关系模型
type Relation struct {
	// TODO: 定义关注关系模型字段
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	ToUserID  uint      `json:"to_user_id" gorm:"index;not null"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 表名
func (Relation) TableName() string {
	return "relations"
}
