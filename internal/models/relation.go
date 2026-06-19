package models

import "time"

// Relation 关注关系模型
type Relation struct {
	ID        int64     `json:"id,string" gorm:"primaryKey"`
	UserID    int64     `json:"user_id,string" gorm:"uniqueIndex:idx_user_follow;not null"`
	ToUserID  int64     `json:"to_user_id,string" gorm:"uniqueIndex:idx_user_follow;not null"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	ToUser    User      `json:"to_user" gorm:"foreignKey:ToUserID"`
}

// TableName 表名
func (Relation) TableName() string {
	return "relations"
}
