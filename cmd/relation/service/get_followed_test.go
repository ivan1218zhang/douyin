package service_test

import (
	"context"
	"douyin/cmd/relation/dal"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"testing"
)

func TestGetFollowed(t *testing.T) {
	dal.Init()
	s := service.NewGetFollowedService(context.Background())
	_, err := s.GetFollowedUser(&relation.MGetFollowReq{UserId: 94053855959977984})
	if err != nil {
		t.Fatal(err)
	}
}
