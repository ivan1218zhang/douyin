package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	c, err := favoriteservice.NewClient(
		"favorite",
		client.WithHostPorts(constants.FavoriteServiceWithHostPorts),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// FavoriteAction 点赞操作
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (*api.FavoriteActionResp, error) {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.FavoriteActionResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
	}, nil
}

// MGetFavorite 喜欢列表
func MGetFavorite(ctx context.Context, req *favorite.MGetFavoriteVideoReq) (*api.MGetFavoriteResp, error) {
	resp, err := favoriteClient.MGetFavoriteVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetFavoriteResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		VideoList:     pack.VideoList(resp.VideoList),
	}, nil
}
