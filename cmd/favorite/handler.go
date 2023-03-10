package main

import (
	"context"
	"douyin/cmd/favorite/pack"
	"douyin/cmd/favorite/service"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// MIsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MIsFavorite(ctx context.Context, req *favorite.MIsFavoriteReq) (resp *favorite.MIsFavoriteResp, err error) {
	resp = new(favorite.MIsFavoriteResp)
	isFavors, err := service.NewMIsFavoriteService(ctx).MIsFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFavoriteList = isFavors
	return
}

// MCountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MCountFavorite(ctx context.Context, req *favorite.MCountFavoriteReq) (resp *favorite.MCountFavoriteResp, err error) {

	resp = new(favorite.MCountFavoriteResp)
	counts, err := service.NewMCountFavoriteService(ctx).MCountFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCountList = counts
	return
}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (resp *favorite.FavoriteActionResp, err error) {
	resp = new(favorite.FavoriteActionResp)
	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetFavoriteVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MGetFavoriteVideo(ctx context.Context, req *favorite.MGetFavoriteVideoReq) (resp *favorite.MGetFavoriteVideoResp, err error) {
	resp = new(favorite.MGetFavoriteVideoResp)
	videos, err := service.NewMGetFavoriteVideoService(ctx).MGetFavoriteVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return
}
