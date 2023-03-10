package service_test

import (
	"context"
	"douyin/cmd/relation/dal"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"testing"
)

func TestRelationAction(t *testing.T) {
	dal.Init()
	// 第一次添加
	err := service.NewRelationActionService(context.Background()).RelationAction(&relation.RelationActionReq{UserId: 1, ToUserId: 2, ActionType: 1})
	if err != nil {
		t.Fatal(err)
	}
	// 第二次重复添加
	err = service.NewRelationActionService(context.Background()).RelationAction(&relation.RelationActionReq{UserId: 1, ToUserId: 2, ActionType: 1})
	// 应添加失败返回错误
	if err == nil {
		t.Fatal(err)
	}
}
