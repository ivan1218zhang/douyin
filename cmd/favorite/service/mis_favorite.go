package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
)

type MIsFavoriteService struct {
	ctx context.Context
}

// NewMIsFavoriteService new MIsFavoriteService
func NewMIsFavoriteService(ctx context.Context) *MIsFavoriteService {
	return &MIsFavoriteService{ctx: ctx}
}

func (s *MIsFavoriteService) MIsFavorite(req *favorite.MIsFavoriteReq) ([]bool, error) {
	isFav, err := db.MIsFavorite(s.ctx, req.UserId, req.VideoIdList)
	return isFav, err
}
