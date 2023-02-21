package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
)

var relationClient relationservice.Client

func initRelation() {
	c, err := relationservice.NewClient(
		"relation",
		client.WithHostPorts(constants.RelationServiceWithHostPorts),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

// GetRelation 获取关注信息
func GetRelation(ctx context.Context, req *relation.GetRelationReq) (*relation.GetRelationResp, error) {
	resp, err := relationClient.GetRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}

// MGetRelation 批量获取关注信息
func MGetRelation(ctx context.Context, req *relation.MGetRelationReq) (*relation.MGetRelationResp, error) {
	resp, err := relationClient.MGetRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
