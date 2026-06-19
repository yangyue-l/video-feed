package utils

import (
	"net/http"
	"video_feed/internal/models"

	"github.com/gin-gonic/gin"
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code:    models.CodeSuccess,
		Message: models.GetMessage(models.CodeSuccess),
		Data:    data,
	})
}

// SuccessWithMessage 成功响应(自定义消息)
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code:    models.CodeSuccess,
		Message: message,
		Data:    data,
	})
}

// SuccessPage 分页成功响应
func SuccessPage(c *gin.Context, data interface{}, total int64, page, size int) {
	c.JSON(http.StatusOK, models.PageResponse{
		Code:    models.CodeSuccess,
		Message: models.GetMessage(models.CodeSuccess),
		Data:    data,
		Total:   total,
		Page:    page,
		Size:    size,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, models.Response{
		Code:    code,
		Message: models.GetMessage(code),
	})
}

// ErrorWithMessage 错误响应(自定义消息)
func ErrorWithMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, models.Response{
		Code:    code,
		Message: message,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, models.Response{
		Code:    models.CodeUnauthorized,
		Message: models.GetMessage(models.CodeUnauthorized),
	})
}

// ParamError 参数错误响应
func ParamError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, models.Response{
		Code:    models.CodeParamError,
		Message: message,
	})
}

// ServerError 服务器错误响应
func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, models.Response{
		Code:    models.CodeServerError,
		Message: models.GetMessage(models.CodeServerError),
	})
}
