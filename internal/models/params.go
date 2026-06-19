package models

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamUserAction 关注/取关用户参数
type ParamUserAction struct {
	ToUserID   int64 `json:"to_user_id" binding:"required"`
	ActionType int   `json:"action_type" binding:"required,oneof=1 2"` // 1-关注 2-取消关注
}

// ParamFavorite 点赞/取消点赞参数
type ParamFavorite struct {
	VideoID    int64 `json:"video_id" binding:"required"`
	ActionType int   `json:"action_type" binding:"required,oneof=1 2"` // 1-点赞 2-取消点赞
}

// ParamCommentAction 评论操作参数
type ParamCommentAction struct {
	VideoID    int64  `json:"video_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
	ActionType int    `json:"action_type" binding:"required,oneof=1 2"` // 1-发布评论 2-删除评论
	CommentID  int64  `json:"comment_id"`                              // 删除评论时需要
}

// ParamVideoList 视频列表查询参数
type ParamVideoList struct {
	UserID     int64 `json:"user_id" form:"user_id"`
	Page       int   `json:"page" form:"page" binding:"required,min=1"`
	Size       int   `json:"size" form:"size" binding:"required,min=1,max=20"`
	LatestTime int64 `json:"latest_time" form:"latest_time"`
}

// ParamVideoFeed 视频流查询参数
type ParamVideoFeed struct {
	LatestTime int64 `json:"latest_time" form:"latest_time"`
}
