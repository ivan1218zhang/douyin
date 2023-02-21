package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

type GetFollowedService struct {
	ctx context.Context
}

func NewGetFollowedService(ctx context.Context) *GetFollowedService {
	return &GetFollowedService{ctx: ctx}
}

func (s *GetFollowedService) GetFollowedUser(req *relation.MGetFollowReq) ([]*common.User, error) {
	res, err := db.GetFollowedUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{IdList: res, UserId: req.Id})
	if err != nil {
		return nil, err
	}
	return resp.UserList, nil
}
