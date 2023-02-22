package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"douyin/pkg/repository"
	"douyin/pkg/util"
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
	videoId := db.Snowflake.NextSnowID()
	videoName := fmt.Sprintf("%d.mp4", videoId)
	coverName := fmt.Sprintf("%d.png", videoId)
	playUrl := constants.CDN.Url + videoName
	coverUrl := constants.CDN.Url + coverName
	// 存入数据库
	videoModel := &repository.Video{
		ID:       videoId,
		AuthorId: req.UserId,
		Title:    req.Title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	// 存到cdn
	go saveVideoCdn(videoName, coverName, req.Data)
	return db.CreateVideo(p.ctx, videoModel)
}

// 把视频和封面存到七牛云
func saveVideoCdn(videoName string, coverName string, data []byte) {
	// 视频存本地
	err := ioutil.WriteFile(constants.CDN.LocalPath+videoName, data, 0644)
	if err != nil {
		panic(err)
	}
	// 截取封面
	err = util.SavePicture(videoName, coverName)
	if err != nil {
		panic(err)
	}
	//视频存七牛云
	err = util.UploadCdn(videoName)
	if err != nil {
		panic(err)
	}
	//封面存七牛云
	err = util.UploadCdn(coverName)
	if err != nil {
		panic(err)
	}
	// TODO 容灾处理
}
