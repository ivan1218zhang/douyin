package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/rpc"
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
	resp, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{IdList: res, UserId: req.Id})
	if err != nil {
		return nil, err
	}
	return resp.UserList, nil
}
