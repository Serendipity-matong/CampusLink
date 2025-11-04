package response

import (
	"github.com/campuslink/campuslink/common/common-core/errors"
)

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
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

// Success 成功响应
func Success(data interface{}) *Response {
	return &Response{
		Code:    errors.CodeSuccess,
		Message: "success",
		Data:    data,
	}
}

// SuccessWithMessage 带消息的成功响应
func SuccessWithMessage(message string, data interface{}) *Response {
	return &Response{
		Code:    errors.CodeSuccess,
		Message: message,
		Data:    data,
	}
}

// Error 错误响应
func Error(err *errors.BizError) *Response {
	return &Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    nil,
	}
}

// ErrorWithMessage 自定义错误响应
func ErrorWithMessage(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// Page 分页成功响应
func Page(data interface{}, total int64, page, size int) *PageResponse {
	return &PageResponse{
		Code:    errors.CodeSuccess,
		Message: "success",
		Data:    data,
		Total:   total,
		Page:    page,
		Size:    size,
	}
}


