package service_test

import (
	"context"
	"douyin/cmd/comment/dal"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
	"fmt"
	"testing"
)

func TestCommentAction(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	service := service.NewCommentActionService(ctx)

	// Test case 1: 发布评论
	req1 := &comment.CommentActionReq{
		UserId:      94401017218600960,
		VideoId:     1,
		ActionType:  1,
		CommentText: "test1111111",
	}
	c, err := service.CommentAction(req1)
	if err != nil {
		t.Errorf("CommentAction returned an error: %v", err)
	}
	fmt.Printf("%+vn\n", c)

	// Test case 2: 删除评论
	req2 := &comment.CommentActionReq{
		ActionType: 2,
		CommentId:  c.Id,
	}
	if _, err := service.CommentAction(req2); err != nil {
		t.Errorf("CommentAction returned an error: %v", err)
	}
}
