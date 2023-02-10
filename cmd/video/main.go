package main

import (
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/conf"
	"douyin/pkg/db"
	"log"
	"os"
)

func main() {
	// init数据库
	err := db.Init()
	if err != nil {
		return
	}
	// 创建存放视频的文件夹
	err = os.MkdirAll(conf.CDN.LocalPath, os.ModePerm)
	if err != nil {
		return
	}

	svr := video.NewServer(new(VideoServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
