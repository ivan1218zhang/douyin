package service

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/pkg/db"
	"fmt"
	"testing"
	"time"
)

func TestMGetPublishService_MGetPublishList(t *testing.T) {
	// 初始化数据库
	err := db.Init()
	if err != nil {
		panic(err)
		return
	}
	// 生成请求
	req := &video.MGetPublishReq{
		UserId: 0,
	}
	// 请求
	ctx := context.Background()
	mGetPublishService := NewMGetPublishService(ctx)
	list, err := mGetPublishService.MGetPublishList(req)
	fmt.Print(list)
	// 等待goroutine结束
	time.Sleep(10000 * time.Millisecond)
}
