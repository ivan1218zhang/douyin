package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/client"
)

var videoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient("video", client.WithHostPorts("0.0.0.0:8888"))
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
