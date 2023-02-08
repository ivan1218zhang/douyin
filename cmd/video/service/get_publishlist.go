package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/video"
)

type GetPublishListService struct {
	ctx context.Context
}

// NewGetPublishListService new GetPublishListService
func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

// QueryNoteService query list of note info
func (s *GetPublishListService) GetPublishListService(req *video.GetPublishListReq) ([]*common.Video, error) {
	videoModels, err := db.QueryVideo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	result := make([]*common.Video, len(videoModels))
	for i, v := range videoModels {
		result[i] = &common.Video{
			Id: v.ID,
			Author: &common.User{
				Id:   v.Author.ID,
				Name: v.Author.UserName,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		}
	}
	return result, nil
}
