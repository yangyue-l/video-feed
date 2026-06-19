package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// UserDao 用户数据访问层
type UserDao struct {
	db *gorm.DB
}

// NewUserDao 创建用户DAO
func NewUserDao() *UserDao {
	return &UserDao{
		db: database.GetDB(),
	}
}

// Create 创建用户
func (d *UserDao) Create(user *models.User) error {
	return d.db.Create(user).Error
}

// FindByUsername 根据用户名查找用户
func (d *UserDao) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找用户
func (d *UserDao) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := d.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (d *UserDao) Update(user *models.User) error {
	return d.db.Save(user).Error
}

// GetFollowCount 获取关注数
func (d *UserDao) GetFollowCount(userID uint) (int64, error) {
	var count int64
	err := d.db.Model(&models.Relation{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

// GetFollowerCount 获取粉丝数
func (d *UserDao) GetFollowerCount(userID uint) (int64, error) {
	var count int64
	err := d.db.Model(&models.Relation{}).Where("to_user_id = ?", userID).Count(&count).Error
	return count, err
}

// IsFollow 检查是否关注
func (d *UserDao) IsFollow(userID, toUserID uint) bool {
	var count int64
	d.db.Model(&models.Relation{}).Where("user_id = ? AND to_user_id = ?", userID, toUserID).Count(&count)
	return count > 0
}
