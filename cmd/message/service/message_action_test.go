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

	req := &message.MassageActionReq{
		//UserId:   99233186235551744,
		//ToUserId: 99229651494244352,
		UserId:     99229651494244352,
		ToUserId:   99233186235551744,
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
