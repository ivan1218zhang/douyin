package main

import (
	favorite "douyin/kitex_gen/favorite/favoriteservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", constants.FavoriteServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	svr := favorite.NewServer(new(FavoriteServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
