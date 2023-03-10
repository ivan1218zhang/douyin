package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
)

type MCountFavoriteService struct {
	ctx context.Context
}

// NewMCountFavoriteService new MCountFavoriteService
func NewMCountFavoriteService(ctx context.Context) *MCountFavoriteService {
	return &MCountFavoriteService{ctx: ctx}
}

// MCountFavorite count favorite number of given video IDs
func (s *MCountFavoriteService) MCountFavorite(req *favorite.MCountFavoriteReq) ([]int64, error) {
	counts, err := db.MCountFavorite(s.ctx, req.VideoIdList)
	return counts, err
}
