package rpc

import (
	"context"
	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/pack"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"time"
)

var commentClient commentservice.Client

func initComment() {
	c, err := commentservice.NewClient(
		"comment",
		client.WithHostPorts(constants.CommentServiceWithHostPorts),
		client.WithRPCTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

// CommentAction 评论操作
func CommentAction(ctx context.Context, req *comment.CommentActionReq) (*api.CommentActionResp, error) {
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.CommentActionResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		Comment:       pack.Comment(resp.Comment),
	}, nil
}

// MGetComment 评论列表
func MGetComment(ctx context.Context, req *comment.MGetCommentReq) (*api.MGetCommentResp, error) {
	resp, err := commentClient.MGetComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return &api.MGetCommentResp{
		StatusCode:    resp.BaseResp.StatusCode,
		StatusMessage: resp.BaseResp.StatusMessage,
		CommentList:   pack.CommentList(resp.CommentList),
	}, nil
}
