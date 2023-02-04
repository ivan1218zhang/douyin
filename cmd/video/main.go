package main

import (
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/db"
	"log"
)

func main() {
	// init数据库
	err := db.Init()
	if err != nil {
		return
	}
	svr := video.NewServer(new(VideoServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
