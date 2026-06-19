package controller

import (
	"video_feed/internal/service"
	"video_feed/internal/utils"

	"github.com/gin-gonic/gin"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	token, err := userService.Register(req.Username, req.Password)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, token)
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	token, err := userService.Login(req.Username, req.Password)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, token)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userID := c.GetUint("user_id")

	userService := service.NewUserService()
	userInfo, err := userService.GetUserInfo(userID)
	if err != nil {
		utils.ErrorWithMessage(c, 1, err.Error())
		return
	}

	utils.Success(c, userInfo)
}
