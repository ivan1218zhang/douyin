package service_test

import (
	"context"
	"douyin/cmd/relation/dal"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func Test_GetRelation(t *testing.T) {
	dal.Init()
	followCount, followerCount, isFollow, _ := service.NewGetRelationService(context.Background()).GetRelation(&relation.GetRelationReq{
		UserId:   2,
		ToUserId: 1,
	})
	fmt.Printf("%d %d %v", followCount, followerCount, isFollow)
}
