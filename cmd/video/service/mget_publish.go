package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/video"
)

type MGetPublishService struct {
	ctx context.Context
}

// NewMGetPublishService new MGetPublishService
func NewMGetPublishService(ctx context.Context) *MGetPublishService {
	return &MGetPublishService{ctx: ctx}
}

// MGetPublishList query list of note info
func (s *MGetPublishService) MGetPublishList(req *video.MGetPublishReq) ([]*common.Video, error) {
	videos, err := db.MGetPublish(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// TODO 从user微服务模块中得到用户信息
	idList := make([]int64, len(videos))
	videoIDList := make([]int64, len(videos))
	for i := 0; i < len(idList); i++ {
		idList[i] = videos[i].AuthorId
		videoIDList[i] = videos[i].Id
	}
	user, err := rpc.GetUser(s.ctx, &user.GetUserReq{
		Id:     req.Id,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	favoriteCount, err := rpc.MCountFavorite(s.ctx, &favorite.MCountFavoriteReq{VideoIdList: videoIDList})
	if err != nil {
		return nil, err
	}
	commentCount, err := rpc.MCountComment(s.ctx, &comment.MCountCommentReq{VideoIdList: videoIDList})
	if err != nil {
		return nil, err
	}
	isfavorites, err := rpc.MIsFavorite(s.ctx, &favorite.MIsFavoriteReq{UserId: req.UserId, VideoIdList: videoIDList})
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(videos); i++ {
		videos[i].Author = user
		videos[i].FavoriteCount = favoriteCount[i]
		videos[i].CommentCount = commentCount[i]
		videos[i].IsFavorite = isfavorites[i]
	}
	return videos, nil
}
