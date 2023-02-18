package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/relation"
)

type MGetFriendService struct {
	ctx context.Context
}

// NewGetFriendService new GetFriendService
func NewGetFriendService(ctx context.Context) *MGetFriendService {
	return &MGetFriendService{ctx: ctx}
}

type void struct {
}

// MGetFriend multiple get list of friend
func (s *MGetFriendService) MGetFriend(req *relation.MGetFriendReq) ([]*common.User, error) {

	var res []*common.User
	// 获取关注用户id
	followedUser, err := db.MGetFollowedUser(s.ctx, req.UserId)
	if err != nil {
		return res, err
	}
	// 获取关注者用户id
	followerUser, err := db.MGetFollowerUser(s.ctx, req.UserId)
	if err != nil {
		return res, err
	}
	var userID []int64
	var member void
	// 取交集
	set := make(map[int64]void)
	for i := range followedUser {
		set[followedUser[i]] = member
	}
	for i := range followedUser {
		_, exists := set[followerUser[i]]
		if exists == true {
			userID = append(userID, followerUser[i])
		}
	}
	res, err = db.MGetUsers(s.ctx, userID)
	if err != nil {
		return res, err
	}
	return res, err

}
