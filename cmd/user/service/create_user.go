package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/repository"
	"douyin/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService 实例化CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user
func (s *CreateUserService) CreateUser(req *user.CreateUserReq) (int64, error) {
	// 因为有查询+插入两步操作 需要加锁保证并发安全
	util.KLock.Lock(req.Username)
	defer util.KLock.Unlock(req.Username)
	// 检查用户名是否重复
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}
	// bcrypt密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	password := string(hash)
	// 插入数据库
	user1 := &repository.User{
		UserName: req.Username,
		Password: password,
	}
	return db.CreateUser(s.ctx, user1)
}
