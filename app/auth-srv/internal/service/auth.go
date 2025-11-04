package service

import (
	"context"

	pb "github.com/campuslink/campuslink/api/auth/v1"
	"github.com/campuslink/campuslink/app/auth-srv/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// AuthService 认证服务
type AuthService struct {
	pb.UnimplementedAuthServer

	authUsecase *biz.AuthUsecase
	log         *log.Helper
}

// NewAuthService 创建认证服务
func NewAuthService(authUsecase *biz.AuthUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		authUsecase: authUsecase,
		log:         log.NewHelper(logger),
	}
}

// Login 用户登录
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	s.log.WithContext(ctx).Infof("Login: %v", req.Username)
	
	token, user, err := s.authUsecase.Login(ctx, req.Username, req.Password, req.LoginType)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
		UserInfo: &pb.UserInfo{
			UserId:    uint64(user.ID),
			Username:  user.Username,
			Phone:     user.Phone,
			StudentId: user.StudentID,
			RealName:  user.RealName,
			School:    user.School,
			Avatar:    user.Avatar,
			Role:      user.Role,
		},
	}, nil
}

// Register 用户注册
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	s.log.WithContext(ctx).Infof("Register: %v", req.Username)
	
	userID, err := s.authUsecase.Register(ctx, &biz.RegisterInfo{
		Username:         req.Username,
		Password:         req.Password,
		Phone:            req.Phone,
		StudentID:        req.StudentId,
		RealName:         req.RealName,
		School:           req.School,
		VerificationCode: req.VerificationCode,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		UserId:  userID,
		Message: "注册成功",
	}, nil
}

// RefreshToken 刷新Token
func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	s.log.WithContext(ctx).Infof("RefreshToken")
	
	token, err := s.authUsecase.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshTokenReply{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

// VerifyToken 验证Token
func (s *AuthService) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenReply, error) {
	s.log.WithContext(ctx).Infof("VerifyToken")
	
	claims, err := s.authUsecase.VerifyToken(ctx, req.Token)
	if err != nil {
		return &pb.VerifyTokenReply{Valid: false}, nil
	}

	return &pb.VerifyTokenReply{
		Valid:    true,
		UserId:   uint64(claims.UserID),
		Username: claims.Username,
		Role:     claims.Role,
	}, nil
}

// Logout 用户登出
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	s.log.WithContext(ctx).Infof("Logout")
	
	err := s.authUsecase.Logout(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutReply{
		Message: "登出成功",
	}, nil
}

// ChangePassword 修改密码
func (s *AuthService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	s.log.WithContext(ctx).Infof("ChangePassword: %v", req.UserId)
	
	err := s.authUsecase.ChangePassword(ctx, uint(req.UserId), req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &pb.ChangePasswordReply{
		Message: "密码修改成功",
	}, nil
}


