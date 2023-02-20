package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
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
	// rpc获取社交信息
	res, err := rpc.GetRelation(s.ctx, &relation.GetRelationReq{UserId: req.UserId, ToUserId: user1.ID})
	if err != nil {
		return nil, err
	}
	u := pack.User(user1)
	u.FollowCount = int64(res.FollowCount)
	u.FollowerCount = int64(res.FollowerCount)
	u.IsFollow = res.IsFollow
	return u, nil
}
