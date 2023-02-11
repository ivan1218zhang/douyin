package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/common"
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
	videoModels, err := db.QueryVideo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(videoModels) == 0 {
		return nil, nil
	}
	userReq := new(user.GetUserReq)
	userReq.Id = req.UserId
	userResp, err := rpc.GetUserById(s.ctx, userReq)
	if err != nil {
		return nil, err
	}
	result := make([]*common.Video, len(videoModels))
	for i, v := range videoModels {
		result[i] = &common.Video{
			Id: v.ID,
			Author: &common.User{
				Id:            userResp.User.Id,
				Name:          userResp.User.Name,
				FollowCount:   userResp.User.FollowCount,
				FollowerCount: userResp.User.FollowerCount,
				IsFollow:      userResp.User.IsFollow,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		}
	}
	return result, nil
}
