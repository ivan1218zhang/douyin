package rpc

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"time"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient(
		"user",
		client.WithHostPorts("0.0.0.0:8889"),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// MGetUser 用户信息
func MGetUser(ctx context.Context, req *user.MGetUserReq) ([]*common.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.UserList, nil
}
