package controller

import (
	"video_feed/internal/models"
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register 用户注册
func Register(c *gin.Context) {
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		utils.ParamError(c, Translate(err))
		return
	}

	if err := service.Register(p.Username, p.Password); err != nil {
		zap.L().Error("service.Register() failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		utils.ParamError(c, Translate(err))
		return
	}

	token, err := service.Login(p.Username, p.Password)
	if err != nil {
		zap.L().Error("service.Login() failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, token)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userID := c.GetInt64("user_id")

	userInfo, err := service.GetUserInfo(userID)
	if err != nil {
		zap.L().Error("service.GetUserInfo() failed", zap.Error(err))
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, userInfo)
}
