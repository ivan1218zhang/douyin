package service

import (
	"context"
	"douyin/cmd/favorite/dal"
	"douyin/kitex_gen/favorite"
	"log"
	"testing"
)

func TestFavoriteActionService_FavoriteAction(t *testing.T) {
	dal.Init()

	req := favorite.FavoriteActionReq{
		UserId:     1,
		VideoId:    1,
		ActionType: 1,
	}
	err := NewFavoriteActionService(context.Background()).FavoriteAction(&req)

	if err != nil {
		log.Fatal(err)
	}

	isFavorReq := favorite.MIsFavoriteReq{
		UserId:      1,
		VideoIdList: []int64{1},
	}

	isFavorArr, err := NewMIsFavoriteService(context.Background()).MIsFavorite(&isFavorReq)

	if err != nil {
		log.Fatal(err)
	}

	if isFavorArr[0] != true {
		log.Fatalln("set favorite failed")
	}

	req = favorite.FavoriteActionReq{
		UserId:     1,
		VideoId:    1,
		ActionType: 2,
	}
	err = NewFavoriteActionService(context.Background()).FavoriteAction(&req)

	if err != nil {
		log.Fatal(err)
	}

	isFavorReq = favorite.MIsFavoriteReq{
		UserId:      1,
		VideoIdList: []int64{1},
	}

	isFavorArr, err = NewMIsFavoriteService(context.Background()).MIsFavorite(&isFavorReq)

	if err != nil {
		log.Fatal(err)
	}

	if isFavorArr[0] != false {
		log.Fatalln("cancel favorite failed")
	}
}
