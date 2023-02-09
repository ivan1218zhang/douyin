package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/video"
)

type MGetVideoService struct {
	ctx context.Context
}

// NewMGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo multiple get list of video info
func (s *MGetVideoService) MGetVideo(req *video.MGetVideoReq) ([]*common.Video, error) {
	videos, err := db.MGetVideo(s.ctx, req.LatestTime)
	if err != nil {
		return nil, err
	}
	// TODO 从user微服务模块中得到用户信息

	return videos, nil
}
