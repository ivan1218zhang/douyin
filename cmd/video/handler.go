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
func (s *VideoServiceImpl) GetVideos(ctx context.Context, getVideosReq *video.GetVideosReq) (resp *video.GetVideosResp, err error) {
	// TODO: Your code here...
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, publishReq *video.PublishReq) (resp *video.PublishResp, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListReq) (resp *video.GetPublishListResp, err error) {
	resp = new(video.GetPublishListResp)
	// TODO:check param
	videos, err := service.NewGetPublishListService(ctx).GetPublishListService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos
	return resp, nil
}
