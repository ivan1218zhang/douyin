package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/client"
	"time"
)

var videoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient(
		"video",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

// MGetVideo 视频流
func MGetVideo(ctx context.Context, req *video.MGetVideoReq) (*api.MGetVideoResp, error) {
	resp, err := videoClient.MGetVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetVideoResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		VideoList:     pack.VideoList(resp.VideoList),
		NextTime:      0,
	}, nil
}

// Publish 上传视频
func Publish(ctx context.Context, req *video.PublishReq) (*api.PublishResp, error) {
	resp, err := videoClient.Publish(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.PublishResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
	}, nil
}

// MGetPublish 用户的投稿
func MGetPublish(ctx context.Context, req *video.MGetPublishReq) (*api.MGetPublishResp, error) {
	resp, err := videoClient.MGetPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetPublishResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		VideoList:     pack.VideoList(resp.VideoList),
	}, nil
}
