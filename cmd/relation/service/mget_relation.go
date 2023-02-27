package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type MGetRelationService struct {
	ctx context.Context
}

// NewMGetRelationService new MGetRelationService
func NewMGetRelationService(ctx context.Context) *MGetRelationService {
	return &MGetRelationService{ctx: ctx}
}

func (g *MGetRelationService) MGetRelation(req *relation.MGetRelationReq) ([]int32, []int32, []bool, error) {
	followCountList, err := db.MCountFollow(g.ctx, req.ToUserIdList)
	if err != nil {
		return make([]int32, len(req.ToUserIdList)), make([]int32, len(req.ToUserIdList)), make([]bool, len(req.ToUserIdList)), err
	}
	followerCountList, err := db.MCountFollower(g.ctx, req.ToUserIdList)
	if err != nil {
		return followCountList, make([]int32, len(req.ToUserIdList)), make([]bool, len(req.ToUserIdList)), err
	}
	if req.UserId == -1 {
		return followCountList, followerCountList, make([]bool, len(req.ToUserIdList)), err
	}
	isFollowList, err := db.MIsFollow(g.ctx, req.UserId, req.ToUserIdList)
	if err != nil {
		return followCountList, followerCountList, make([]bool, len(req.ToUserIdList)), err
	}
	return followCountList, followerCountList, isFollowList, nil
}
