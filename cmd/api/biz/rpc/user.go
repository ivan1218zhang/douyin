package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"strconv"
	"time"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient(
		"user",
		client.WithHostPorts("0.0.0.0:8889"),
		client.WithRPCTimeout(30*time.Second),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser 用户注册
func CreateUser(ctx context.Context, req *user.CreateUserReq) (*api.RegisterResp, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.RegisterResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserID:        resp.UserId,
		Token:         strconv.FormatInt(resp.UserId, 10),
	}, nil
}

// CheckUser 用户登录
func CheckUser(ctx context.Context, req *user.CheckUserReq) (*api.LoginResp, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.LoginResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserID:        resp.UserId,
		Token:         strconv.FormatInt(resp.UserId, 10),
	}, nil
}

// GetUser 用户信息
func GetUser(ctx context.Context, req *user.GetUserReq) (*api.GetUserResp, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.GetUserResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		User:          pack.User(resp.User),
	}, nil
}
