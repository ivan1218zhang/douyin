package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/message"
	"testing"
)

func TestMessageActionService(t *testing.T) {

	// 初始化数据库
	db.Init()

	// 生成请求

	req := &message.MessageActionReq{
		UserId:     1,
		ToUserId:   2,
		ActionType: 3,
		Content:    "打法",
	}
	// 请求
	ctx := context.Background()
	err := NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		println(err)
	}

}
