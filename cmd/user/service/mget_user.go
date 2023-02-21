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
	// rpc获取社交信息

	userList := pack.Users(modelUsers)

	res, err := rpc.MGetRelation(s.ctx, &relation.MGetRelationReq{UserId: req.UserId, ToUserIdList: req.IdList})
	if err != nil {
		return nil, err
	}

	for i, u := range userList {
		u.FollowCount = int64(res.FollowCountList[i])
		u.FollowerCount = int64(res.FollowerCountList[i])
		u.IsFollow = res.IsFollowList[i]
	}

	return userList, nil
}
