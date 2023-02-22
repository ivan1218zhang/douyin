package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type GetRelationService struct {
	ctx context.Context
}

// NewGetRelationService new GetRelationService
func NewGetRelationService(ctx context.Context) *GetRelationService {
	return &GetRelationService{ctx: ctx}
}

func (g *GetRelationService) GetRelation(req *relation.GetRelationReq) (int32, int32, bool, error) {
	followCount, err := db.CountFollow(g.ctx, req.ToUserId)
	if err != nil {
		return 0, 0, false, err
	}
	followerCount, err := db.CountFollower(g.ctx, req.ToUserId)
	if err != nil {
		return followCount, 0, false, err
	}
	isFollow, err := db.IsFollow(g.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return followCount, followerCount, false, err
	}
	return followCount, followerCount, isFollow, nil
}
