package main

import (
	"douyin/cmd/favorite/dal"
	"douyin/cmd/favorite/rpc"
	favorite "douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", constants.FavoriteServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	Init()

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr), // address
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//server.WithMuxTransport(),                                          // Multiplex
		//server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
