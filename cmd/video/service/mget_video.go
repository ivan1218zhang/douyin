package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/video"
	"time"
)

type MGetVideoService struct {
	ctx context.Context
}

// NewMGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo multiple get list of video info
func (s *MGetVideoService) MGetVideo(req *video.MGetVideoReq) ([]*common.Video, int64, error) {
	// 视频
	videos, err := db.MGetVideo(s.ctx, req.LatestTime)
	if err != nil {
		return nil, time.Now().UnixMilli(), err
	}
	// 获取next_time
	id := videos[len(videos)-1].Id
	nextTime, err := db.GetVideoCreatedAt(s.ctx, id)
	// TODO 从user微服务模块中得到用户信息

	return videos, nextTime, nil
}
