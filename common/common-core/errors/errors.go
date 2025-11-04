package errors

import (
	"fmt"
	"net/http"
)

// BizError 业务错误
type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *BizError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// 常用错误码定义
const (
	// 通用错误码 (1xxx)
	CodeSuccess           = 0
	CodeInternalError     = 1000
	CodeInvalidParams     = 1001
	CodeNotFound          = 1002
	CodeAlreadyExists     = 1003
	CodePermissionDenied  = 1004
	CodeUnauthorized      = 1005
	CodeTooManyRequests   = 1006

	// 用户相关错误码 (2xxx)
	CodeUserNotFound      = 2001
	CodeUserExists        = 2002
	CodePasswordIncorrect = 2003
	CodeTokenInvalid      = 2004
	CodeTokenExpired      = 2005
	CodePhoneExists       = 2006
	CodeStudentNotVerified = 2007

	// 商品相关错误码 (3xxx)
	CodeProductNotFound   = 3001
	CodeProductOffline    = 3002
	CodeInsufficientStock = 3003
	CodeProductExpired    = 3004

	// 任务相关错误码 (4xxx)
	CodeTaskNotFound      = 4001
	CodeTaskClosed        = 4002
	CodeTaskAccepted      = 4003

	// 订单相关错误码 (5xxx)
	CodeOrderNotFound     = 5001
	CodeOrderCanceled     = 5002
	CodeOrderPaid         = 5003
	CodeOrderCompleted    = 5004

	// 支付相关错误码 (6xxx)
	CodePaymentFailed     = 6001
	CodeRefundFailed      = 6002
	CodeInsufficientBalance = 6003
)

// 预定义错误
var (
	ErrInternalError     = NewError(CodeInternalError, "内部服务错误")
	ErrInvalidParams     = NewError(CodeInvalidParams, "参数错误")
	ErrNotFound          = NewError(CodeNotFound, "资源不存在")
	ErrAlreadyExists     = NewError(CodeAlreadyExists, "资源已存在")
	ErrPermissionDenied  = NewError(CodePermissionDenied, "权限不足")
	ErrUnauthorized      = NewError(CodeUnauthorized, "未授权")
	ErrTooManyRequests   = NewError(CodeTooManyRequests, "请求过于频繁")

	ErrUserNotFound      = NewError(CodeUserNotFound, "用户不存在")
	ErrUserExists        = NewError(CodeUserExists, "用户已存在")
	ErrPasswordIncorrect = NewError(CodePasswordIncorrect, "密码错误")
	ErrTokenInvalid      = NewError(CodeTokenInvalid, "Token无效")
	ErrTokenExpired      = NewError(CodeTokenExpired, "Token已过期")

	ErrProductNotFound   = NewError(CodeProductNotFound, "商品不存在")
	ErrProductOffline    = NewError(CodeProductOffline, "商品已下架")
	ErrInsufficientStock = NewError(CodeInsufficientStock, "库存不足")

	ErrTaskNotFound      = NewError(CodeTaskNotFound, "任务不存在")
	ErrOrderNotFound     = NewError(CodeOrderNotFound, "订单不存在")
	ErrPaymentFailed     = NewError(CodePaymentFailed, "支付失败")
)

// NewError 创建业务错误
func NewError(code int, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

// WithDetails 添加错误详情
func (e *BizError) WithDetails(details string) *BizError {
	return &BizError{
		Code:    e.Code,
		Message: e.Message,
		Details: details,
	}
}

// HTTPStatus 获取对应的 HTTP 状态码
func (e *BizError) HTTPStatus() int {
	switch e.Code {
	case CodeNotFound, CodeUserNotFound, CodeProductNotFound, CodeTaskNotFound, CodeOrderNotFound:
		return http.StatusNotFound
	case CodeUnauthorized, CodeTokenInvalid, CodeTokenExpired:
		return http.StatusUnauthorized
	case CodePermissionDenied:
		return http.StatusForbidden
	case CodeInvalidParams:
		return http.StatusBadRequest
	case CodeTooManyRequests:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}


