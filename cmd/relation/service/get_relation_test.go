package service

import (
	"context"
	"douyin/cmd/relation/dal"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func Test_GetRelation(t *testing.T) {
	dal.Init()
	followCount, followerCount, isFollow, _ := NewGetRelationService(context.Background()).GetRelation(&relation.GetRelationReq{
		UserId:   2,
		ToUserId: 1,
	})
	fmt.Printf("%d %d %v", followCount, followerCount, isFollow)
}
