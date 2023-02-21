package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/constants"
	"time"

	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient(
		"user",
		client.WithHostPorts(constants.UserServiceWithHostPorts),
		client.WithRPCTimeout(30*time.Second),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// MGetUser 用户信息
func MGetUser(ctx context.Context, req *user.MGetUserReq) (*user.MGetUserResp, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &user.MGetUserResp{
		BaseResp: resp.BaseResp,
		UserList: resp.UserList,
	}, nil
}
