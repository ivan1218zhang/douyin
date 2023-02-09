package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/video"
)

type MGetPublishService struct {
	ctx context.Context
}

// NewMGetPublishService new MGetPublishService
func NewMGetPublishService(ctx context.Context) *MGetPublishService {
	return &MGetPublishService{ctx: ctx}
}

// MGetPublishList query list of note info
func (s *MGetPublishService) MGetPublishList(req *video.MGetPublishReq) ([]*common.Video, error) {
	videos, err := db.QueryVideo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// TODO 从user微服务模块中得到用户信息

	return videos, nil
}
