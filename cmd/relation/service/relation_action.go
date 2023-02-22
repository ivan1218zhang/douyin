package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (r *RelationActionService) RelationAction(req *relation.RelationActionReq) error {
	var err error
	if req.ActionType == 1 {
		var isfollow bool
		isfollow, err = db.IsFollow(r.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
		if !isfollow {
			err = db.AddRelation(r.ctx, req.UserId, req.ToUserId)
		}
	}
	if req.ActionType == 2 {
		err = db.DeleteRelation(r.ctx, req.UserId, req.ToUserId)
	}
	return err
}
