package service_test

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func TestMGetFriendService(t *testing.T) {

	// 初始化数据库
	db.Init()

	// 生成请求
	req := &relation.MGetFriendReq{
		UserId: 99229651494244352,
	}
	// 请求
	ctx := context.Background()
	resp, err := service.NewGetFriendService(ctx).MGetFriend(req)
	if err != nil {
		println(err)
	}
	for i := range resp {
		fmt.Println(*resp[i])
	}
}
