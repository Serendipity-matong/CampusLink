package biz

import (
	"context"
	"time"

	"github.com/campuslink/campuslink/common/common-auth"
	"github.com/campuslink/campuslink/common/common-core/errors"
	"github.com/campuslink/campuslink/common/common-core/utils"
	"github.com/go-kratos/kratos/v2/log"
)

// User 用户业务对象
type User struct {
	ID        uint
	Username  string
	Password  string
	Phone     string
	StudentID string
	RealName  string
	School    string
	Avatar    string
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Token 令牌信息
type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

// RegisterInfo 注册信息
type RegisterInfo struct {
	Username         string
	Password         string
	Phone            string
	StudentID        string
	RealName         string
	School           string
	VerificationCode string
}

// UserRepo 用户仓储接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (uint64, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	GetUserByStudentID(ctx context.Context, studentID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, userID uint, password string) error
}

// AuthUsecase 认证用例
type AuthUsecase struct {
	repo UserRepo
	jwt  *auth.JWT
	log  *log.Helper
}

// NewAuthUsecase 创建认证用例
func NewAuthUsecase(repo UserRepo, jwt *auth.JWT, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		jwt:  jwt,
		log:  log.NewHelper(logger),
	}
}

// Login 用户登录
func (uc *AuthUsecase) Login(ctx context.Context, username, password, loginType string) (*Token, *User, error) {
	var user *User
	var err error

	// 根据登录类型获取用户
	switch loginType {
	case "phone":
		user, err = uc.repo.GetUserByPhone(ctx, username)
	case "student_id":
		user, err = uc.repo.GetUserByStudentID(ctx, username)
	default:
		user, err = uc.repo.GetUserByUsername(ctx, username)
	}

	if err != nil {
		return nil, nil, errors.ErrUserNotFound
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		return nil, nil, errors.ErrPasswordIncorrect
	}

	// 检查用户状态
	if user.Status == "banned" {
		return nil, nil, errors.NewError(errors.CodePermissionDenied, "用户已被封禁")
	}

	// 生成Token
	accessToken, err := uc.jwt.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, nil, errors.ErrInternalError.WithDetails(err.Error())
	}

	refreshToken, err := uc.jwt.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, nil, errors.ErrInternalError.WithDetails(err.Error())
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    7200, // 2小时
	}, user, nil
}

// Register 用户注册
func (uc *AuthUsecase) Register(ctx context.Context, info *RegisterInfo) (uint64, error) {
	// 检查用户名是否已存在
	existUser, _ := uc.repo.GetUserByUsername(ctx, info.Username)
	if existUser != nil {
		return 0, errors.ErrUserExists
	}

	// 检查手机号是否已存在
	if info.Phone != "" {
		existUser, _ = uc.repo.GetUserByPhone(ctx, info.Phone)
		if existUser != nil {
			return 0, errors.NewError(errors.CodePhoneExists, "手机号已存在")
		}
	}

	// 哈希密码
	hashedPassword, err := utils.HashPassword(info.Password)
	if err != nil {
		return 0, errors.ErrInternalError.WithDetails(err.Error())
	}

	// 创建用户
	user := &User{
		Username:  info.Username,
		Password:  hashedPassword,
		Phone:     info.Phone,
		StudentID: info.StudentID,
		RealName:  info.RealName,
		School:    info.School,
		Role:      "student",
		Status:    "active",
	}

	userID, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, errors.ErrInternalError.WithDetails(err.Error())
	}

	return userID, nil
}

// RefreshToken 刷新Token
func (uc *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	// 解析Token
	claims, err := uc.jwt.ParseToken(refreshToken)
	if err != nil {
		return nil, errors.ErrTokenInvalid
	}

	// 生成新Token
	accessToken, err := uc.jwt.GenerateToken(claims.UserID, claims.Username, claims.Role)
	if err != nil {
		return nil, errors.ErrInternalError.WithDetails(err.Error())
	}

	newRefreshToken, err := uc.jwt.GenerateToken(claims.UserID, claims.Username, claims.Role)
	if err != nil {
		return nil, errors.ErrInternalError.WithDetails(err.Error())
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    7200,
	}, nil
}

// VerifyToken 验证Token
func (uc *AuthUsecase) VerifyToken(ctx context.Context, token string) (*auth.Claims, error) {
	claims, err := uc.jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// Logout 用户登出
func (uc *AuthUsecase) Logout(ctx context.Context, token string) error {
	// TODO: 将token加入黑名单（使用Redis）
	return nil
}

// ChangePassword 修改密码
func (uc *AuthUsecase) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	// 获取用户
	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return errors.ErrUserNotFound
	}

	// 验证旧密码
	if !utils.CheckPassword(oldPassword, user.Password) {
		return errors.ErrPasswordIncorrect
	}

	// 哈希新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.ErrInternalError.WithDetails(err.Error())
	}

	// 更新密码
	err = uc.repo.UpdatePassword(ctx, userID, hashedPassword)
	if err != nil {
		return errors.ErrInternalError.WithDetails(err.Error())
	}

	return nil
}


