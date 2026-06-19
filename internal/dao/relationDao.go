package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// RelationDao 关注关系数据访问层
type RelationDao struct {
	db *gorm.DB
}

// NewRelationDao 创建关注关系DAO
func NewRelationDao() *RelationDao {
	return &RelationDao{
		db: database.GetDB(),
	}
}

// Create 创建关注关系
func (d *RelationDao) Create(relation *models.Relation) error {
	return d.db.Create(relation).Error
}

// Delete 删除关注关系
func (d *RelationDao) Delete(userID, toUserID uint) error {
	return d.db.Where("user_id = ? AND to_user_id = ?", userID, toUserID).Delete(&models.Relation{}).Error
}

// IsFollow 检查是否已关注
func (d *RelationDao) IsFollow(userID, toUserID uint) bool {
	var count int64
	d.db.Model(&models.Relation{}).Where("user_id = ? AND to_user_id = ?", userID, toUserID).Count(&count)
	return count > 0
}

// GetFollowList 获取关注列表
func (d *RelationDao) GetFollowList(userID uint) ([]models.User, error) {
	var users []models.User
	err := d.db.Joins("JOIN relations ON relations.to_user_id = users.id").
		Where("relations.user_id = ?", userID).
		Order("relations.created_at DESC").
		Find(&users).Error
	return users, err
}

// GetFollowerList 获取粉丝列表
func (d *RelationDao) GetFollowerList(userID uint) ([]models.User, error) {
	var users []models.User
	err := d.db.Joins("JOIN relations ON relations.user_id = users.id").
		Where("relations.to_user_id = ?", userID).
		Order("relations.created_at DESC").
		Find(&users).Error
	return users, err
}
