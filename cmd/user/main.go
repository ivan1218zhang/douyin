package main

import (
	user2 "douyin/cmd/user"
	user "douyin/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(user2.UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
