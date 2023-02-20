package service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/cmd/favorite/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/video"
)

type MGetFavoriteVideoService struct {
	ctx context.Context
}

// NewMGetFavoriteVideoService new MGetFavoriteVideoService
func NewMGetFavoriteVideoService(ctx context.Context) *MGetFavoriteVideoService {
	return &MGetFavoriteVideoService{ctx: ctx}
}

// MGetFavoriteVideo gets users favorite videos
func (s *MGetFavoriteVideoService) MGetFavoriteVideo(req *favorite.MGetFavoriteVideoReq) ([]*common.Video, error) {
	records, err := db.MGetFavoriteByUserID(s.ctx, req.UserId)
	vIds := make([]int64, len(records))
	for i := 0; i < len(records); i++ {
		vIds[i] = records[i].VideoId
	}
	resp, err := rpc.MGetVideoById(s.ctx, &video.MGetVideoByIdReq{
		IdList: vIds,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return resp.VideoList, err
}
