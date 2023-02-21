package main

import (
	"context"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionReq) (resp *relation.RelationActionResp, err error) {
	// TODO: Your code here...
	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// MGetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFollow(ctx context.Context, req *relation.MGetFollowReq) (resp *relation.MGetFollowResp, err error) {
	users, err := service.NewGetFollowedService(ctx).GetFollowedUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.SetUserList(users)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// MGetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFollower(ctx context.Context, req *relation.MGetFollowerReq) (resp *relation.MGetFollowerResp, err error) {
	users, err := service.NewGetFollowerService(ctx).GetFollowerUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.SetUserList(users)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// MGetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFriend(ctx context.Context, req *relation.MGetFriendReq) (resp *relation.MGetFriendResp, err error) {
	// TODO: Your code here...
	resp = new(relation.MGetFriendResp)
	users, err := service.NewGetFriendService(ctx).MGetFriend(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserList = users
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = users
	return resp, err
}

// GetRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetRelation(ctx context.Context, req *relation.GetRelationReq) (resp *relation.GetRelationResp, err error) {
	// TODO: Your code here...
	resp = new(relation.GetRelationResp)
	followCount, followerCount, isFollow, err := service.NewGetRelationService(ctx).GetRelation(req)
	resp.FollowCount = followCount
	resp.FollowerCount = followerCount
	resp.IsFollow = isFollow
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// MGetRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetRelation(ctx context.Context, req *relation.MGetRelationReq) (resp *relation.MGetRelationResp, err error) {
	// TODO: Your code here...
	resp = new(relation.MGetRelationResp)
	followCountList, followerCountList, isFollowList, err := service.NewMGetRelationService(ctx).MGetRelation(req)
	resp.FollowCountList = followCountList
	resp.FollowerCountList = followerCountList
	resp.IsFollowList = isFollowList
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}
