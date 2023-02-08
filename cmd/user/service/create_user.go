package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	ctx context.Context
}

// 实例化CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserReq) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	password := string(hash)
	user := &repository.User{
		UserName: req.Username,
		Password: password,
	}
	return db.CreateUser(s.ctx, user)
}
