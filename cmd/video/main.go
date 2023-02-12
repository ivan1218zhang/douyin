package main

import (
	"douyin/cmd/video/dal"
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/conf"
	"log"
	"os"
)

func Init() {
	// init数据库
	dal.Init()
	// 创建存放视频的文件夹
	_ = os.MkdirAll(conf.CDN.LocalPath, os.ModePerm)
}

func main() {
	Init()
	svr := video.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
