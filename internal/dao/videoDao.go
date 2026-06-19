package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// VideoDao 视频数据访问层
type VideoDao struct {
	db *gorm.DB
}

// NewVideoDao 创建视频DAO
func NewVideoDao() *VideoDao {
	return &VideoDao{
		db: database.GetDB(),
	}
}

// Create 创建视频
func (d *VideoDao) Create(video *models.Video) error {
	return d.db.Create(video).Error
}

// FindByID 根据ID查找视频
func (d *VideoDao) FindByID(id uint) (*models.Video, error) {
	var video models.Video
	err := d.db.First(&video, id).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

// GetFeed 获取视频Feed流
func (d *VideoDao) GetFeed(latestTime int64, limit int) ([]models.Video, error) {
	var videos []models.Video
	query := d.db.Order("created_at DESC")
	if latestTime > 0 {
		query = query.Where("created_at < ?", latestTime)
	}
	err := query.Limit(limit).Find(&videos).Error
	return videos, err
}

// GetUserVideos 获取用户发布的视频
func (d *VideoDao) GetUserVideos(userID uint) ([]models.Video, error) {
	var videos []models.Video
	err := d.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&videos).Error
	return videos, err
}

// UpdateFavoriteCount 更新点赞数
func (d *VideoDao) UpdateFavoriteCount(videoID uint, delta int64) error {
	return d.db.Model(&models.Video{}).Where("id = ?", videoID).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", delta)).Error
}

// UpdateCommentCount 更新评论数
func (d *VideoDao) UpdateCommentCount(videoID uint, delta int64) error {
	return d.db.Model(&models.Video{}).Where("id = ?", videoID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", delta)).Error
}
