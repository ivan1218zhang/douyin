package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction req.ActionType == 1 set favorite , 2 del favorite
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionReq) error {
	var err error
	if req.ActionType == 1 {
		err = db.CreateFavorite(s.ctx, req.UserId, req.VideoId)
	} else {
		err = db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return err
}
