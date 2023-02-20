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
	"time"
)

type MGetVideoService struct {
	ctx context.Context
}

// NewMGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo multiple get list of video info
func (s *MGetVideoService) MGetVideo(req *video.MGetVideoReq) ([]*common.Video, int64, error) {
	// 视频
	videos, err := db.MGetVideo(s.ctx, req.LatestTime)
	if err != nil {
		return nil, time.Now().UnixMilli(), err
	}
	// 获取next_time
	nextTime := time.Now().UnixMilli()
	if len(videos) > 1 {
		id := videos[len(videos)-1].Id
		nextTime, err := db.GetVideoCreatedAt(s.ctx, id)
		if err != nil {
			return nil, nextTime, err
		}
	}
	// TODO 从user微服务模块中得到用户信息
	idList := make([]int64, len(videos))
	for i := 0; i < len(idList); i++ {
		idList[i] = videos[i].AuthorId
	}
	users, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{
		IdList: idList,
		UserId: 1,
	})
	if err != nil {
		return nil, time.Now().UnixMilli(), err
	}
	for i := 0; i < len(videos); i++ {
		videos[i].Author = users[i]
	}
	return videos, nextTime, nil
}

// MGetVideoById multiple get list of video info by id
func (s *MGetVideoService) MGetVideoById(req *video.MGetVideoByIdReq) ([]*common.Video, error) {
	// 视频
	videos, err := db.MGetVideoById(s.ctx, req.IdList)
	if err != nil {
		return nil, err
	}

	idList := make([]int64, len(videos))
	videoIDList := make([]int64, len(videos))
	for i := 0; i < len(idList); i++ {
		idList[i] = videos[i].AuthorId
		videoIDList[i] = videos[i].Id
	}
	users, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{
		IdList: idList,
		UserId: 0,
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
		videos[i].Author = users[i]
		videos[i].FavoriteCount = favoriteCount[i]
		videos[i].CommentCount = commentCount[i]
		videos[i].IsFavorite = isfavorites[i]
	}
	return videos, nil
}
