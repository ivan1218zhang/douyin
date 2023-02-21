package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

type GetFollowerService struct {
	ctx context.Context
}

func NewGetFollowerService(ctx context.Context) *GetFollowerService {
	return &GetFollowerService{ctx: ctx}
}

func (s *GetFollowerService) GetFollowerUser(req *relation.MGetFollowerReq) ([]*common.User, error) {
	res, err := db.GetFollowerUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userList, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{IdList: res, UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return userList, nil
}
