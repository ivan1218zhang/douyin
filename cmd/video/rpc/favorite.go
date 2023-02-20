package rpc

import (
	"context"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
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

// MCountFavorite 喜欢数量
func MCountFavorite(ctx context.Context, req *favorite.MCountFavoriteReq) ([]int64, error) {
	resp, err := favoriteClient.MCountFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.FavoriteCountList, err
}
