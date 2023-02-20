package main

import (
	"douyin/cmd/comment/rpc"
	comment "douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	rpc.Init()
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	svr := comment.NewServer(new(CommentServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
