package data

import (
	"github.com/campuslink/campuslink/app/auth-srv/internal/conf"
	"github.com/campuslink/campuslink/common/common-auth"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewJWT, NewUserRepo)

// Data 数据层
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData 创建数据层
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

// NewDB 创建数据库连接
func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(c.Database.ConnMaxLifetime)
	
	// 自动迁移
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	
	return db
}

// NewRedis 创建Redis连接
func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           c.Redis.DB,
		DialTimeout:  c.Redis.DialTimeout,
		ReadTimeout:  c.Redis.ReadTimeout,
		WriteTimeout: c.Redis.WriteTimeout,
	})
	return rdb
}

// NewJWT 创建JWT管理器
func NewJWT(c *conf.Data) *auth.JWT {
	return auth.NewJWT(&auth.JWTConfig{
		Secret:     c.JWT.Secret,
		Issuer:     c.JWT.Issuer,
		Expiration: c.JWT.Expiration,
	})
}

