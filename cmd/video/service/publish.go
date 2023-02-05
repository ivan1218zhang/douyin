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
	videoName := fmt.Sprintf("%d.mp4", videoId)
	coverName := fmt.Sprintf("%d.png", videoId)
	playUrl := fmt.Sprintf("%s%s", conf.CDN.Url, videoName)
	coverUrl := fmt.Sprintf("%s%s", conf.CDN.Url, coverName)
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
	go saveVideoCdn(videoName, coverName, req.Data)
	return db.CreateVideo(p.ctx, videoModel)
}

// 把视频和封面存到七牛云
func saveVideoCdn(videoName string, coverName string, data []byte) error {
	// 视频存本地
	err := ioutil.WriteFile(fmt.Sprintf("./public/%s", videoName), data, 0644)
	if err != nil {
		return nil
	}
	return nil
}
