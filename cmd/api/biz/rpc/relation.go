package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/client"
	"time"
)

var relationClient relationservice.Client

func initRelation() {
	c, err := relationservice.NewClient(
		"relation",
		client.WithHostPorts("0.0.0.0:8892"),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

// RelationAction 关注操作
func RelationAction(ctx context.Context, req *relation.RelationActionReq) (*api.RelationActionResp, error) {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.RelationActionResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
	}, nil
}

// MGetFollow 关注列表
func MGetFollow(ctx context.Context, req *relation.MGetFollowReq) (*api.MGetFollowResp, error) {
	resp, err := relationClient.MGetFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetFollowResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserList:      pack.UserList(resp.UserList),
	}, nil
}

// MGetFollower 粉丝列表
func MGetFollower(ctx context.Context, req *relation.MGetFollowerReq) (*api.MGetFollowerResp, error) {
	resp, err := relationClient.MGetFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetFollowerResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		UserList:      pack.UserList(resp.UserList),
	}, nil
}
