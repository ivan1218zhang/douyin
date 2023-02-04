package main

import (
	"douyin/cmd/user/dal"
	user "douyin/kitex_gen/user/userservice"
	"douyin/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {

	dal.Init()
	addr, _ := net.ResolveTCPAddr("tcp", conf.UserPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := user.NewServer(new(UserServiceImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
