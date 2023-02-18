package service

import (
	"context"
	"douyin/cmd/api/biz/rpc"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/video"
	"log"
	"sync"
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
	videoModels, err := db.MGetVideo(s.ctx, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}
	videos := pack.Videos(videoModels)
	wg := sync.WaitGroup{}
	log.Print(len(videos))

	for i := range videos {

		wg.Add(1)
		go func(index int, v *common.Video) {
			defer wg.Done()
			if v == nil {
				return
			}
			resp, err := rpc.GetUser(s.ctx, &user.GetUserReq{Id: v.AuthorId})
			if err != nil {
				log.Print(err)
				return
			}

			v.Author = &common.User{}
			u := resp.User
			v.Author.Id = u.ID
			v.Author.Name = u.Name
			v.Author.FollowerCount = u.FollowerCount
			v.Author.FollowCount = u.FollowCount

		}(i, videos[i])
	}

	/*

		wg.Add(2)
		go func() {
			defer wg.Done()
			uIds := pack.UserIds(videos)
			users, err := rpc.MGetUser(s.ctx, &user.MGetUserReq{IdList: uIds})
			if err != nil {
				log.Print(err)
			}

			for i := range videos {
				videos[i].Author = users[i]
			}
		}()

	*/
	wg.Wait()
	return videos, videoModels[len(videos)-1].CreatedAt.Unix() - 1, nil
}
