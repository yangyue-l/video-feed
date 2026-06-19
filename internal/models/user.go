package models

import "time"

// User 用户模型
type User struct {
	// TODO: 定义用户模型字段
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;size:32;not null"`
	Password  string    `json:"-" gorm:"size:128;not null"`
	Nickname  string    `json:"nickname" gorm:"size:32"`
	Avatar    string    `json:"avatar" gorm:"size:256"`
	Signature string    `json:"signature" gorm:"size:256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse 用户响应(不包含敏感信息)
type UserResponse struct {
	// TODO: 定义用户响应字段
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Signature     string `json:"signature"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
