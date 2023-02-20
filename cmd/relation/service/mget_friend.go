package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

type MGetFriendService struct {
	ctx context.Context
}

// NewGetFriendService new GetFriendService
func NewGetFriendService(ctx context.Context) *MGetFriendService {
	return &MGetFriendService{ctx: ctx}
}

// MGetFriend multiple get list of friend
func (s *MGetFriendService) MGetFriend(req *relation.MGetFriendReq) ([]*common.User, error) {
	followers, err := db.MGetFollowerUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	follows, err := db.MGetFollowedUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	set := map[int64]interface{}{}
	for i := 0; i < len(follows); i++ {
		set[follows[i]] = 0
	}
	res := make([]int64, 0)
	for i := 0; i < len(followers); i++ {
		if set[followers[i]] != nil {
			res = append(res, followers[i])
		}
	}
	resp, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{
		IdList: res,
		UserId: req.UserId,
	})
	return resp.UserList, err
}
