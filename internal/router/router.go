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

		// 需要认证的接口
		authGroup := v1Group.Group("")
		authGroup.Use(middleware.JWTAuth())
		{
			//获取用户信息
			authGroup.GET("/user", controller.GetUserInfo)

			// TODO: 实现视频相关接口
			videoGroup := authGroup.Group("/video")
			{
				videoGroup.POST("/publish", controller.PublishVideo)
				videoGroup.GET("/feed", controller.GetVideoFeed)
				videoGroup.GET("/list", controller.GetUserVideos)
				videoGroup.GET("/:id", controller.GetVideoDetail)
			}

			// TODO: 实现互动相关接口
			interactGroup := authGroup.Group("/interact")
			{
				interactGroup.POST("/favorite", controller.FavoriteVideo)
				interactGroup.GET("/favorite/list", controller.GetFavoriteList)
				interactGroup.POST("/comment", controller.AddComment)
				interactGroup.GET("/comment/list", controller.GetCommentList)
				interactGroup.DELETE("/comment/:id", controller.DeleteComment)
			}

			// TODO: 实现社交相关接口
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
