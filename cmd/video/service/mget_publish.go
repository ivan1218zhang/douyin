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
	userReq := new(user.MGetUserReq)
	for i, v := range videoModels {
		userReq.IdList[i] = v.AuthorId
	}
	usersResp, err := rpc.MGetUser(s.ctx, userReq)
	if err != nil {
		return nil, err
	}
	result := make([]*common.Video, len(videoModels))
	for i, v := range videoModels {
		result[i] = &common.Video{
			Id: v.Id,
			Author: &common.User{
				Id:            usersResp.UserList[i].Id,
				Name:          usersResp.UserList[i].Name,
				FollowCount:   usersResp.UserList[i].FollowCount,
				FollowerCount: usersResp.UserList[i].FollowerCount,
				IsFollow:      usersResp.UserList[i].IsFollow,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		}
	}
	return result, nil
}
