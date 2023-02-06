package main

import (
	"context"

	"douyin/cmd/video/pack"
	"douyin/cmd/video/service"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, req *video.GetVideosReq) (resp *video.GetVideosResp, err error) {
	resp = new(video.GetVideosResp)

	videos, err := service.NewMGetVideoService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, publishReq *video.PublishReq) (resp *video.PublishResp, err error) {
	// TODO: Your code here...
	publishService := service.NewPublishService(ctx)
	err = publishService.Publish(publishReq)
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, getPublishListReq *video.GetPublishListReq) (resp *video.GetPublishListResp, err error) {
	// TODO: Your code here...
	return
}
