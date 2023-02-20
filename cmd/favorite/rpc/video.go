package rpc

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/client"
)

var videoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient(
		"video",
		client.WithHostPorts(constants.VideoServiceWithHostPorts),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func MGetVideoById(ctx context.Context, req *video.MGetVideoByIdReq) (*video.MGetVideoByIdResp, error) {
	resp, err := videoClient.MGetVideoById(ctx, req)
	if err != nil {
		return nil, err
	}
	return &video.MGetVideoByIdResp{
		BaseResp:  resp.BaseResp,
		VideoList: resp.VideoList,
	}, nil
}
