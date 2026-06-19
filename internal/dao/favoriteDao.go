package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// FavoriteDao 点赞数据访问层
type FavoriteDao struct {
	db *gorm.DB
}

// NewFavoriteDao 创建点赞DAO
func NewFavoriteDao() *FavoriteDao {
	return &FavoriteDao{
		db: database.GetDB(),
	}
}

// Create 创建点赞
func (d *FavoriteDao) Create(favorite *models.Favorite) error {
	return d.db.Create(favorite).Error
}

// Delete 删除点赞
func (d *FavoriteDao) Delete(userID, videoID uint) error {
	return d.db.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&models.Favorite{}).Error
}

// IsFavorite 检查是否已点赞
func (d *FavoriteDao) IsFavorite(userID, videoID uint) bool {
	var count int64
	d.db.Model(&models.Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Count(&count)
	return count > 0
}

// GetFavoriteList 获取用户点赞列表
func (d *FavoriteDao) GetFavoriteList(userID uint) ([]models.Video, error) {
	var videos []models.Video
	err := d.db.Joins("JOIN favorites ON favorites.video_id = videos.id").
		Where("favorites.user_id = ?", userID).
		Order("favorites.created_at DESC").
		Find(&videos).Error
	return videos, err
}
