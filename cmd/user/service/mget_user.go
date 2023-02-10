package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
)

type GetMUsersService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) MGetUser(req *user.MGetUserReq) (user []*common.User, err error) {
	user1, err := db.MGetUser(s.ctx, req.IdList)
	if err != nil {
		return nil, err
	}
	return user1, nil
}
