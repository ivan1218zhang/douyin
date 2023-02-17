package main

import (
	"douyin/cmd/message/dal"
	message "douyin/kitex_gen/message/messageservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))
	dal.Init()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
