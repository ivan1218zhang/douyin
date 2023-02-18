package service

import (
	"context"
	"douyin/cmd/favorite/dal/rdb"
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
	if req.ActionType == 1 {
		err := rdb.SetFavorite(req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	} else {
		err := rdb.DelFavorite(req.UserId, req.VideoId)
		if err != nil {
			return err
		}
	}

	return nil
}
