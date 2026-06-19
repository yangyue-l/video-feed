package dao

import (
	"video_feed/internal/database"
	"video_feed/internal/models"

	"gorm.io/gorm"
)

// CommentDao 评论数据访问层
type CommentDao struct {
	db *gorm.DB
}

// NewCommentDao 创建评论DAO
func NewCommentDao() *CommentDao {
	return &CommentDao{
		db: database.GetDB(),
	}
}

// Create 创建评论
func (d *CommentDao) Create(comment *models.Comment) error {
	return d.db.Create(comment).Error
}

// Delete 删除评论
func (d *CommentDao) Delete(id uint) error {
	return d.db.Delete(&models.Comment{}, id).Error
}

// FindByID 根据ID查找评论
func (d *CommentDao) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := d.db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentList 获取视频评论列表
func (d *CommentDao) GetCommentList(videoID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := d.db.Where("video_id = ?", videoID).Order("created_at DESC").Find(&comments).Error
	return comments, err
}
