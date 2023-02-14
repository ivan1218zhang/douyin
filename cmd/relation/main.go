package main

import (
	"douyin/cmd/relation/dal"
	relation "douyin/kitex_gen/relation/relationservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationServiceImpl))
	dal.Init()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
