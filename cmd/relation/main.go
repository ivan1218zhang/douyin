package main

import (
	"douyin/cmd/relation/dal"
	relation "douyin/kitex_gen/relation/relationservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", constants.RelationServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	svr := relation.NewServer(new(RelationServiceImpl), server.WithServiceAddr(addr))
	dal.Init()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
