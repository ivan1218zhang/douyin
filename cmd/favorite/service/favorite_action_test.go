package service

import (
	"context"
	"douyin/cmd/favorite/dal"
	"douyin/kitex_gen/favorite"
	"log"
	"testing"
)

// test all favorite api
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
	req = favorite.FavoriteActionReq{
		UserId:     1,
		VideoId:    3,
		ActionType: 1,
	}
	err = NewFavoriteActionService(context.Background()).FavoriteAction(&req)
	if err != nil {
		log.Fatal(err)
	}
	req = favorite.FavoriteActionReq{
		UserId:     2,
		VideoId:    1,
		ActionType: 1,
	}
	err = NewFavoriteActionService(context.Background()).FavoriteAction(&req)
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
		log.Fatalln("test isFavorite failed")
	}

	countFavorReq := favorite.MCountFavoriteReq{
		VideoIdList: []int64{1},
	}
	favorCountArr, err := NewMCountFavoriteService(context.Background()).MCountFavorite(&countFavorReq)
	if err != nil {
		log.Fatal(err)
	}
	if favorCountArr[0] != 2 {
		log.Fatalln("test countFavorite failed")
	}

	favorVideoReq := favorite.MGetFavoriteVideoReq{
		UserId: 1,
	}
	favorVideoIds, err := NewMGetFavoriteVideoService(context.Background()).MGetFavoriteVideo(&favorVideoReq)
	if err != nil {
		log.Fatal(err)
	}
	if len(favorVideoIds) != 2 {
		log.Fatalln("test favorVideo failed")
	}
	/*-----------------------------------------------------------*/

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

	countFavorReq = favorite.MCountFavoriteReq{
		VideoIdList: []int64{1},
	}
	favorCountArr, err = NewMCountFavoriteService(context.Background()).MCountFavorite(&countFavorReq)
	if err != nil {
		log.Fatal(err)
	}
	if favorCountArr[0] != 1 {
		log.Fatalln("test countFavorite failed")
	}

	favorVideoIds, err = NewMGetFavoriteVideoService(context.Background()).MGetFavoriteVideo(&favorVideoReq)
	if err != nil {
		log.Fatal(err)
	}
	if favorVideoIds[0].FavoriteCount != 3 {
		log.Fatalln("test favorVideo failed")
	}

}
