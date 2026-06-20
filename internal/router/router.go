package router

import (
	"video_feed/config"
	"video_feed/internal/controller"
	"video_feed/internal/logger"
	"video_feed/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(cfg *config.Config) *gin.Engine {
	// 设置运行模式
	gin.SetMode(cfg.Server.Mode)

	r := gin.New()

	// 全局中间件
	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(middleware.CORS())

	// 静态文件服务
	r.Static("/static", "./static")

	// API v1
	v1Group := r.Group("/api/v1")
	{
		// 用户相关(无需认证)
		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/register", controller.Register)
			userGroup.POST("/login", controller.Login)
		}

		// 公开接口（无需登录）
		videoGroup := v1Group.Group("/video")
		{
			videoGroup.GET("/feed", controller.GetVideoFeed)        // 视频流
			videoGroup.GET("/list", controller.GetUserVideos)       // 用户视频列表
			videoGroup.GET("/detail/:id", controller.GetVideoDetail) // 视频详情
		}

		// 评论列表（公开，未登录也能看）
		v1Group.GET("/interact/comment/list", controller.GetCommentList)

		// 需要认证的接口
		authGroup := v1Group.Group("")
		authGroup.Use(middleware.JWTAuth())
		{
			// 获取用户信息
			authGroup.GET("/user", controller.GetUserInfo)

			// 视频发布（需要登录）
			authGroup.POST("/video/publish", controller.PublishVideo)

			// 互动相关接口（需要登录）
			interactGroup := authGroup.Group("/interact")
			{
				interactGroup.POST("/favorite", controller.FavoriteVideo)
				interactGroup.GET("/favorite/list", controller.GetFavoriteList)
				interactGroup.POST("/comment", controller.AddComment)
				interactGroup.DELETE("/comment/:id", controller.DeleteComment)
			}

			// 社交相关接口（需要登录）
			socialGroup := authGroup.Group("/social")
			{
				socialGroup.POST("/follow", controller.FollowUser)
				socialGroup.GET("/follow/list", controller.GetFollowList)
				socialGroup.GET("/follower/list", controller.GetFollowerList)
			}
		}
	}

	return r
}
