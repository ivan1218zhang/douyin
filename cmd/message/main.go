package main

import (
	"douyin/cmd/message/dal"
	message "douyin/kitex_gen/message/messageservice"
	"douyin/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", constants.MessageServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	svr := message.NewServer(new(MessageServiceImpl), server.WithServiceAddr(addr))
	dal.Init()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
