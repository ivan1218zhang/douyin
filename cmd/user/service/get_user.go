package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *user.GetUserReq) (user *common.User, err error) {
	user1, err := db.GetUser(s.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return pack.User(user1), nil
}
