package main

import (
	"context"
	"douyin/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// IsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// MIsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MIsFavorite(ctx context.Context, req *favorite.MIsFavoriteReq) (resp *favorite.MIsFavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// CountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountFavorite(ctx context.Context, req *favorite.CountFavoriteReq) (resp *favorite.CountFavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// MCountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MCountFavorite(ctx context.Context, req *favorite.MCountFavoriteReq) (resp *favorite.MCountFavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (resp *favorite.FavoriteActionResp, err error) {
	// TODO: Your code here...
	return
}

// MGetFavoriteVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MGetFavoriteVideo(ctx context.Context, req *favorite.MGetFavoriteVideoReq) (resp *favorite.MGetFavoriteVideoResp, err error) {
	// TODO: Your code here...
	return
}
