package main

import (
	"douyin/cmd/video/internal/dal"
	video "douyin/kitex_gen/video/video/videoservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
)

func Init() {
	dal.Init()
}

func main() {
	//r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	//if err != nil {
	//	panic(err)
	//}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
