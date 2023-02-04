package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/conf"
	"douyin/pkg/repository"
	"fmt"
	"io/ioutil"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

func (p *PublishService) Publish(req *video.PublishReq) error {
	// 生成视频信息
	videoId := 0
	playUrl := fmt.Sprintf("%s%d.mp4", conf.CDN.Url, videoId)
	coverUrl := fmt.Sprintf("%s%d.png", conf.CDN.Url, videoId)
	// 存入数据库
	videoModel := &repository.Video{
		ID:            0,
		AuthorId:      req.UserId,
		Title:         req.Title,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	// 存到cdn
	go saveVideoCdn(playUrl, coverUrl, req.Data)
	return db.CreateVideo(p.ctx, videoModel)
}

// 把视频和封面存到七牛云
func saveVideoCdn(playUrl string, coverUrl string, data []int8) {
	// 视频存本地
	byteData := []byte{}
	ioutil.WriteFile(playUrl, byteData, 0644)
}
