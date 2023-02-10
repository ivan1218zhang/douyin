package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *user.MGetUserReq) (user []*common.User, err error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.IdList)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
