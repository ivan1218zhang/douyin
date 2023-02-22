package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
)

var commentClient commentservice.Client

func initComment() {
	c, err := commentservice.NewClient(
		"comment",
		client.WithHostPorts(constants.CommentServiceWithHostPorts),
		client.WithRPCTimeout(10*time.Second),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

// MGetComment 评论列表
func MCountComment(ctx context.Context, req *comment.MCountCommentReq) ([]int64, error) {
	resp, err := commentClient.MCountComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.CommentCountList, nil
}
