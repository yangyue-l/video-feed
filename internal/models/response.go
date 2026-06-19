package models

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageResponse 分页响应
type PageResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

// TokenResponse Token响应
type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// 响应状态码
const (
	CodeSuccess       = 0
	CodeError         = 1
	CodeUnauthorized  = 401
	CodeForbidden     = 403
	CodeNotFound      = 404
	CodeParamError    = 400
	CodeServerError   = 500
)

// 响应消息
var messageMap = map[int]string{
	CodeSuccess:      "success",
	CodeError:        "操作失败",
	CodeUnauthorized: "未授权",
	CodeForbidden:    "禁止访问",
	CodeNotFound:     "资源不存在",
	CodeParamError:   "参数错误",
	CodeServerError:  "服务器内部错误",
}

// GetMessage 获取响应消息
func GetMessage(code int) string {
	if msg, ok := messageMap[code]; ok {
		return msg
	}
	return "未知错误"
}
