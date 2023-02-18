package service

import (
	"context"
	"douyin/cmd/favorite/dal/rdb"
	"douyin/kitex_gen/favorite"
)

type MGetFavoriteVideoService struct {
	ctx context.Context
}

// NewMGetFavoriteVideoService new MGetFavoriteVideoService
func NewMGetFavoriteVideoService(ctx context.Context) *MGetFavoriteVideoService {
	return &MGetFavoriteVideoService{ctx: ctx}
}

// MGetFavoriteVideo gets users favorite videos
func (s *MGetFavoriteVideoService) MGetFavoriteVideo(req *favorite.MGetFavoriteVideoReq) ([]int64, error) {
	vIDs, err := rdb.MGetFavoriteVideoID(req.UserId)

	return vIDs, err
}
