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

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.MGetVideoReq) (resp *video.MGetVideoResp, err error) {
	resp = new(video.MGetVideoResp)

	videos, nextTime, err := service.NewMGetVideoService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = nextTime
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishReq) (resp *video.PublishResp, err error) {
	resp = new(video.PublishResp)
	publishService := service.NewPublishService(ctx)
	err = publishService.Publish(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetPublish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetPublish(ctx context.Context, req *video.MGetPublishReq) (resp *video.MGetPublishResp, err error) {
	resp = new(video.MGetPublishResp)
	// TODO:check param
	videos, err := service.NewMGetPublishService(ctx).MGetPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if len(videos) > 0 {

	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// MGetVideoById implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideoById(ctx context.Context, req *video.MGetVideoByIdReq) (resp *video.MGetVideoByIdResp, err error) {
	// TODO: Your code here...
	resp = new(video.MGetVideoByIdResp)
	videos, err := service.NewMGetVideoService(ctx).MGetVideoById(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
