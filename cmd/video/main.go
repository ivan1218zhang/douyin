package main

import (
	"douyin/cmd/video/dal"
	"douyin/cmd/video/mw"
	"douyin/cmd/video/rpc"
	video "douyin/kitex_gen/video/videoservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"os"
)

func Init() {
	// init数据库
	dal.Init()
	// 创建存放视频的文件夹
	_ = os.MkdirAll(constants.CDN.LocalPath, os.ModePerm)
	rpc.Init()
	err := mw.InitRabbitmq()
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	addr, err := net.ResolveTCPAddr("tcp", constants.VideoServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	svr := video.NewServer(new(VideoServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
