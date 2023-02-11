package service

import (
	"context"

	"douyin/cmd/video/internal/dal/db"
	"douyin/cmd/video/internal/pack"

	"douyin/kitex_gen/video/video"
)

type MGetVideoService struct {
	ctx context.Context
}

// NewMGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo multiple get list of video info
func (s *MGetVideoService) MGetVideo(req *video.GetVideosReq) ([]*common.Video, error) {
	videoModels, err := db.MGetVideos(s.ctx, req.LatestTime)
	if err != nil {
		return nil, err
	}
	videos := pack.Videos(videoModels)

	return videos, nil
}
