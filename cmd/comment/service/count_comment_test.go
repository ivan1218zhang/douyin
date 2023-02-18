package service_test

import (
	"context"
	"douyin/cmd/comment/dal"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
	"fmt"
	"testing"
	"time"
)

func TestCountComment(t *testing.T) {
	dal.Init()
	ctx := context.Background()

	service := service.NewCountCommentService(ctx)

	req := &comment.CountCommentReq{
		VideoId: 1,
	}

	start := time.Now()
	count, err := service.CountComment(req)
	if err != nil {
		t.Error(err)
	}
	elapsed := time.Since(start)
	fmt.Println("评论数量：", count)
	fmt.Println("首次执行完成耗时：", elapsed)

	start = time.Now()
	count, err = service.CountComment(req)
	if err != nil {
		t.Error(err)
	}
	elapsed = time.Since(start)
	fmt.Println("评论数量：", count)
	fmt.Println("第二次执行完成耗时：", elapsed)

}
