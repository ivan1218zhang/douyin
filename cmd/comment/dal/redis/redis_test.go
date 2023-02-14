package redis_test

import (
	"context"
	"douyin/cmd/comment/dal"
	"douyin/cmd/comment/dal/redis"
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	redis.Init()
}

func TestGetCommentsForVideo(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	videoId := 1
	start := time.Now()
	_, err := redis.GetCommentsForVideo(ctx, int64(videoId))
	if err != nil {
		t.Fatalf("GetCommentsForVideo return an error: %v", err)
	}
	elapsed := time.Since(start)
	fmt.Println("首次执行完成耗时：", elapsed)

	start = time.Now()
	_, err = redis.GetCommentsForVideo(ctx, int64(videoId))
	if err != nil {
		t.Fatalf("GetCommentsForVideo return an error: %v", err)
	}
	elapsed = time.Since(start)
	fmt.Println("第二次执行完成耗时：", elapsed)
}
