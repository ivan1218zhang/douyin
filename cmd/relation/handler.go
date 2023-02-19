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
	return
}

// MGetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFollow(ctx context.Context, req *relation.MGetFollowReq) (resp *relation.MGetFollowResp, err error) {
	// TODO: Your code here...
	return
}

// MGetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFollower(ctx context.Context, req *relation.MGetFollowerReq) (resp *relation.MGetFollowerResp, err error) {
	// TODO: Your code here...
	return
}

// CountFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollow(ctx context.Context, req *relation.CountFollowReq) (resp *relation.CountFollowResp, err error) {
	// TODO: Your code here...
	return
}

// CountFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollower(ctx context.Context, req *relation.CountFollowerReq) (resp *relation.CountFollowerResp, err error) {
	// TODO: Your code here...
	return
}

// IsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollow(ctx context.Context, req *relation.IsFollowReq) (resp *relation.IsFollowResp, err error) {
	// TODO: Your code here...
	return
}

// MCountFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MCountFollow(ctx context.Context, req *relation.MCountFollowReq) (resp *relation.MCountFollowResp, err error) {
	// TODO: Your code here...
	return
}

// MCountFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MCountFollower(ctx context.Context, req *relation.MCountFollowerReq) (resp *relation.MCountFollowerResp, err error) {
	// TODO: Your code here...
	return
}

// MIsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MIsFollow(ctx context.Context, req *relation.MIsFollowReq) (resp *relation.MIsFollowResp, err error) {
	// TODO: Your code here...
	return
}

// MGetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetFriend(ctx context.Context, req *relation.MGetFriendReq) (resp *relation.MGetFriendResp, err error) {
	// TODO: Your code here...
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
	return
}

// MGetRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MGetRelation(ctx context.Context, req *relation.MGetRelationReq) (resp *relation.MGetRelationResp, err error) {
	// TODO: Your code here...
	return
}
