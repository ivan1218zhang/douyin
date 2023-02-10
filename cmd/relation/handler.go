package main

import (
	"context"
	"douyin/kitex_gen/relation"
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
