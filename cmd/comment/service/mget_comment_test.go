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

func TestMGetComment(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	service := service.NewMGetCommentService(ctx)

	req1 := &comment.MGetCommentReq{
		VideoId: 1,
	}

	start := time.Now()
	_, err := service.MGetComment(req1)
	if err != nil {
		t.Fatalf("MGetComment return an error: %v", err)
	}
	elapsed := time.Since(start)
	fmt.Println("首次执行完成耗时：", elapsed)

	start = time.Now()
	_, err = service.MGetComment(req1)
	if err != nil {
		t.Fatalf("MGetComment return an error: %v", err)
	}
	elapsed = time.Since(start)
	fmt.Println("第二次执行完成耗时：", elapsed)
}
