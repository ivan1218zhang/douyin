package service

import (
	"context"
	"douyin/cmd/relation/dal"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func Test_MGetRelation(t *testing.T) {
	dal.Init()
	followCount, followerCount, isFollow, _ := NewMGetRelationService(context.Background()).MGetRelation(&relation.MGetRelationReq{
		UserId:       2,
		ToUserIdList: []int64{1, 3},
	})
	fmt.Printf("%v %v %v", followCount, followerCount, isFollow)
}
