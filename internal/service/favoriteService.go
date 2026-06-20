package service

import (
	"errors"
	"video_feed/internal/dao"
	"video_feed/internal/database"
	"video_feed/internal/models"
	"video_feed/internal/utils"

	"gorm.io/gorm"
)

const (
	Like          = 1
	UnLike        = 2
	LikeSuccess   = "点赞成功"
	UnLikeSuccess = "取消点赞成功"
)

// FavoriteVideo 点赞/取消点赞视频
func FavoriteVideo(userID, videoID int64, actionType int) (msg string, err error) {
	// 1. 参数校验
	if actionType != Like && actionType != UnLike {
		return "", errors.New("无效的操作类型")
	}

	// 2. 使用数据库事务
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查是否已点赞
		var count int64
		if err := tx.Model(&models.Favorite{}).
			Where("user_id = ? AND video_id = ?", userID, videoID).
			Count(&count).Error; err != nil {
			return errors.New("查询点赞状态失败")
		}
		likeFlag := count > 0

		if actionType == Like {
			// 点赞
			if likeFlag {
				return errors.New("您已经点赞")
			}
			// 创建点赞记录
			f := &models.Favorite{
				ID:      utils.GenerateID(),
				UserID:  userID,
				VideoID: videoID,
			}
			if err := tx.Create(f).Error; err != nil {
				return errors.New("点赞失败")
			}
			// 更新点赞计数
			if err := tx.Model(&models.Video{}).Where("id = ?", videoID).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return errors.New("更新点赞数失败")
			}
			msg = LikeSuccess
		} else {
			// 取消点赞
			if !likeFlag {
				return errors.New("您还没有点赞")
			}
			// 删除点赞记录
			if err := tx.Where("user_id = ? AND video_id = ?", userID, videoID).
				Delete(&models.Favorite{}).Error; err != nil {
				return errors.New("取消点赞失败")
			}
			// 更新点赞计数
			if err := tx.Model(&models.Video{}).Where("id = ?", videoID).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", -1)).Error; err != nil {
				return errors.New("更新点赞数失败")
			}
			msg = UnLikeSuccess
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return msg, nil
}

// GetFavoriteList 获取用户点赞列表
func GetFavoriteList(userID int64) ([]models.VideoResponse, error) {
	videoList, err := dao.GetFavoriteList(userID)
	if err != nil {
		return nil, err
	}
	var videoRes []models.VideoResponse

	for _, video := range videoList {
		//获取用户信息
		author, err := dao.FindUserByVideoID(video.ID)
		if err != nil {
			continue
		}
		authorRes := &models.UserResponse{
			ID:       author.ID,
			Username: author.Username,
		}

		oVideo := &models.VideoResponse{
			ID:            video.ID,
			Author:        *authorRes,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
		videoRes = append(videoRes, *oVideo)
	}

	return videoRes, nil
}
