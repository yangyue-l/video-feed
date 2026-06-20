package service

import (
	"errors"
	"mime/multipart"
	"video_feed/internal/dao"
	"video_feed/internal/models"
	"video_feed/internal/utils"
)

// PublishVideo 发布视频
func PublishVideo(userID int64, title string, videoFile *multipart.FileHeader, coverFile *multipart.FileHeader) error {
	// 1. 检查OSS客户端是否初始化
	if utils.GetOSSClient() == nil {
		return errors.New("文件上传服务未初始化")
	}

	// 2. 校验文件类型
	if !utils.IsVideoFile(videoFile.Filename) {
		return errors.New("不支持的视频格式")
	}
	if !utils.IsImageFile(coverFile.Filename) {
		return errors.New("不支持的封面图片格式")
	}

	// 3. 上传视频到OSS
	playURL, err := utils.GetOSSClient().UploadVideo(videoFile)
	if err != nil {
		return errors.New("视频上传失败: " + err.Error())
	}

	// 4. 上传封面到OSS
	coverURL, err := utils.GetOSSClient().UploadCover(coverFile)
	if err != nil {
		return errors.New("封面上传失败: " + err.Error())
	}

	// 5. 生成视频ID并创建视频记录
	video := &models.Video{
		ID:       utils.GenerateID(),
		UserID:   userID,
		Title:    title,
		PlayURL:  playURL,
		CoverURL: coverURL,
	}

	// 6. 保存到数据库
	if err := dao.CreateVideo(video); err != nil {
		return errors.New("保存视频信息失败")
	}

	return nil
}

// GetVideoFeed 获取视频Feed流
func GetVideoFeed(userID int64, latestTime int64, limit int) ([]models.VideoResponse, error) {
	// 1. 获取视频列表
	videos, err := dao.GetVideoFeed(latestTime, limit)
	if err != nil {
		return nil, err
	}

	// 2. 转换为响应格式
	var videoList []models.VideoResponse
	for _, video := range videos {
		// 获取视频作者信息
		author, err := dao.FindUserByID(video.UserID)
		if err != nil {
			continue // 跳过找不到作者的视频
		}

		// 构建作者响应
		authorResp := models.UserResponse{
			ID:            author.ID,
			Username:      author.Username,
			Nickname:      author.Nickname,
			Avatar:        author.Avatar,
			Signature:     author.Signature,
			FollowCount:   0,     // TODO: 查询关注数
			FollowerCount: 0,     // TODO: 查询粉丝数
			IsFollow:      false, // TODO: 查询是否关注
		}

		// 检查当前用户是否点赞了该视频
		isFavorite := false
		if userID > 0 {
			isFavorite = dao.IsFavorite(userID, video.ID)
		}

		videoResp := models.VideoResponse{
			ID:            video.ID,
			Author:        authorResp,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, videoResp)
	}

	return videoList, nil
}

// GetUserVideos 获取用户视频列表
func GetUserVideos(userID int64) ([]models.VideoResponse, error) {
	// TODO: 实现获取用户视频列表逻辑

	//1. 获取用户信息
	user, err := dao.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("查找用户失败")
	}
	//1.2 构建用户返回信息
	userRes := &models.UserResponse{
		ID:            userID,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Avatar:        user.Avatar,
		Signature:     user.Signature,
		FollowCount:   0,     // TODO: 查询关注数
		FollowerCount: 0,     // TODO: 查询粉丝数
		IsFollow:      false, // TODO: 查询是否关注

	}

	//2. 获取视频列表
	videos, err := dao.GetUserVideos(userID)
	if err != nil {
		return nil, errors.New("获取用户视频失败")
	}

	var videoList []models.VideoResponse

	for _, video := range videos {
		isFavorite := false
		if userID > 0 {
			isFavorite = dao.IsFavorite(userID, video.ID)
		}

		videoResp := models.VideoResponse{
			ID:            video.ID,
			Author:        *userRes,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, videoResp)
	}

	return videoList, nil
}

// GetVideoDetail 获取视频详情
func GetVideoDetail(userID, videoID int64) (*models.VideoResponse, error) {
	// 1. 获取视频信息
	video, err := dao.FindVideoByID(videoID)
	if err != nil {
		return nil, errors.New("视频不存在")
	}

	// 2. 获取视频作者信息
	author, err := dao.FindUserByID(video.UserID)
	if err != nil {
		return nil, errors.New("获取作者信息失败")
	}

	// 3. 构建作者响应
	authorResp := models.UserResponse{
		ID:            author.ID,
		Username:      author.Username,
		Nickname:      author.Nickname,
		Avatar:        author.Avatar,
		Signature:     author.Signature,
		FollowCount:   0,     // TODO: 查询关注数
		FollowerCount: 0,     // TODO: 查询粉丝数
		IsFollow:      false, // TODO: 查询是否关注
	}

	// 4. 检查当前用户是否点赞了该视频
	isFavorite := false
	if userID > 0 {
		isFavorite = dao.IsFavorite(userID, video.ID)
	}

	// 5. 构建视频响应
	videoResp := &models.VideoResponse{
		ID:            video.ID,
		Author:        authorResp,
		PlayURL:       video.PlayURL,
		CoverURL:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}

	return videoResp, nil
}
