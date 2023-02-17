package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/message"
	"fmt"
	"testing"
)

func TestGetMessageListService(t *testing.T) {

	// 初始化数据库
	db.Init()

	// 生成请求

	req := &message.MessageChatReq{
		UserId:     1,
		FromUserId: 2,
	}
	// 请求
	ctx := context.Background()
	res, err := NewGetMessageChatService(ctx).GetMessageChat(req)
	if err != nil {
		println(err)
	}
	for i := range res {
		fmt.Println(res[i])
	}
}
