package main

import (
	"douyin/cmd/user/dal"
	"douyin/cmd/user/rpc"
	user "douyin/kitex_gen/user/userservice"
	"douyin/pkg/constants"
	"douyin/pkg/util"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func Init() {
	dal.Init()
	rpc.Init()
	util.KLock = util.NewKeyLock()
}

func main() {
	//r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	//if err != nil {
	//	panic(err)
	//}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceWithHostPorts)
	if err != nil {
		panic(err)
	}
	Init()

	svr := user.NewServer(new(UserServiceImpl),
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
