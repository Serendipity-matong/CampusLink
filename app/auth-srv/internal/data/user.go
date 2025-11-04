package data

import (
	"context"
	"time"

	"github.com/campuslink/campuslink/app/auth-srv/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// User 用户持久化对象
type User struct {
	ID        uint           `gorm:"primarykey"`
	Username  string         `gorm:"uniqueIndex;size:50;not null"`
	Password  string         `gorm:"size:255;not null"`
	Phone     string         `gorm:"uniqueIndex;size:20"`
	StudentID string         `gorm:"index;size:50"`
	RealName  string         `gorm:"size:50"`
	School    string         `gorm:"size:100"`
	Avatar    string         `gorm:"size:255"`
	Role      string         `gorm:"size:20;default:student"`
	Status    string         `gorm:"size:20;default:active"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

// userRepo 用户仓储实现
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 创建用户仓储
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (uint64, error) {
	po := &User{
		Username:  user.Username,
		Password:  user.Password,
		Phone:     user.Phone,
		StudentID: user.StudentID,
		RealName:  user.RealName,
		School:    user.School,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Status:    user.Status,
	}
	
	if err := r.data.db.WithContext(ctx).Create(po).Error; err != nil {
		return 0, err
	}
	
	return uint64(po.ID), nil
}

// GetUserByID 根据ID获取用户
func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*biz.User, error) {
	var po User
	if err := r.data.db.WithContext(ctx).Where("id = ?", id).First(&po).Error; err != nil {
		return nil, err
	}
	return r.poToBiz(&po), nil
}

// GetUserByUsername 根据用户名获取用户
func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	var po User
	if err := r.data.db.WithContext(ctx).Where("username = ?", username).First(&po).Error; err != nil {
		return nil, err
	}
	return r.poToBiz(&po), nil
}

// GetUserByPhone 根据手机号获取用户
func (r *userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	var po User
	if err := r.data.db.WithContext(ctx).Where("phone = ?", phone).First(&po).Error; err != nil {
		return nil, err
	}
	return r.poToBiz(&po), nil
}

// GetUserByStudentID 根据学号获取用户
func (r *userRepo) GetUserByStudentID(ctx context.Context, studentID string) (*biz.User, error) {
	var po User
	if err := r.data.db.WithContext(ctx).Where("student_id = ?", studentID).First(&po).Error; err != nil {
		return nil, err
	}
	return r.poToBiz(&po), nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	po := &User{
		ID:        user.ID,
		Username:  user.Username,
		Phone:     user.Phone,
		StudentID: user.StudentID,
		RealName:  user.RealName,
		School:    user.School,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Status:    user.Status,
	}
	return r.data.db.WithContext(ctx).Save(po).Error
}

// UpdatePassword 更新密码
func (r *userRepo) UpdatePassword(ctx context.Context, userID uint, password string) error {
	return r.data.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", userID).
		Update("password", password).Error
}

// poToBiz 转换PO到业务对象
func (r *userRepo) poToBiz(po *User) *biz.User {
	return &biz.User{
		ID:        po.ID,
		Username:  po.Username,
		Password:  po.Password,
		Phone:     po.Phone,
		StudentID: po.StudentID,
		RealName:  po.RealName,
		School:    po.School,
		Avatar:    po.Avatar,
		Role:      po.Role,
		Status:    po.Status,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}
}


