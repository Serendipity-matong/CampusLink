package service

import (
	"context"

	pb "github.com/campuslink/campuslink/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

// UserService 用户服务
type UserService struct {
	pb.UnimplementedUserServer
	log *log.Helper
}

// NewUserService 创建用户服务
func NewUserService(logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper(logger),
	}
}

// GetUser 获取用户信息
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	s.log.WithContext(ctx).Infof("GetUser: %v", req.UserId)
	
	return &pb.GetUserReply{
		User: &pb.UserInfo{
			UserId:   req.UserId,
			Username: "demo_user",
			Nickname: "演示用户",
			Phone:    "13800138000",
			Role:     "student",
			Verified: true,
		},
	}, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	s.log.WithContext(ctx).Infof("UpdateUser: %v", req.UserId)
	return &pb.UpdateUserReply{Message: "更新成功"}, nil
}

// ListUsers 获取用户列表
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	s.log.WithContext(ctx).Infof("ListUsers")
	return &pb.ListUsersReply{
		Users: []*pb.UserInfo{},
		Total: 0,
		Page:  req.Page,
		PageSize: req.PageSize,
	}, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	s.log.WithContext(ctx).Infof("DeleteUser: %v", req.UserId)
	return &pb.DeleteUserReply{Message: "删除成功"}, nil
}

// VerifyStudent 学生认证
func (s *UserService) VerifyStudent(ctx context.Context, req *pb.VerifyStudentRequest) (*pb.VerifyStudentReply, error) {
	s.log.WithContext(ctx).Infof("VerifyStudent: %v", req.UserId)
	return &pb.VerifyStudentReply{
		Message:  "认证成功",
		Verified: true,
	}, nil
}

// GetUserStats 获取用户统计
func (s *UserService) GetUserStats(ctx context.Context, req *pb.GetUserStatsRequest) (*pb.GetUserStatsReply, error) {
	s.log.WithContext(ctx).Infof("GetUserStats: %v", req.UserId)
	return &pb.GetUserStatsReply{
		OrderCount:   0,
		TaskCount:    0,
		ProductCount: 0,
		CreditScore:  100,
		Balance:      0,
	}, nil
}

// SetUserRole 设置用户角色
func (s *UserService) SetUserRole(ctx context.Context, req *pb.SetUserRoleRequest) (*pb.SetUserRoleReply, error) {
	s.log.WithContext(ctx).Infof("SetUserRole: %v to %v", req.UserId, req.Role)
	return &pb.SetUserRoleReply{Message: "角色设置成功"}, nil
}


