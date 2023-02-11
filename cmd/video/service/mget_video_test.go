package service

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/pkg/db"
	"log"
	"testing"
	"time"
)

func TestMGetVideoService_MGetVideo(t *testing.T) {

	/*
		v := repository.Video{
			AuthorId: 1,
			PlayUrl:  "http://cdn.nobugnolife.com/8_53AF93CE860FF6C54935A98DE33B0548.mp4",
			Title:    "烧",
		}
		for i := 1; i <= 30; i++ {
			v.ID = int64(i)
			db.DB.Create(&v)
		}
	*/
	db.Init()
	/*
		v := repository.Video{}
		db.DB.First(&v, 1)
		log.Print(v.CreatedAt)

	*/
	req := video.MGetVideoReq{
		LatestTime: time.Now().AddDate(0, 0, 0).Unix(),
	}
	vs, err := NewMGetVideoService(context.Background()).MGetVideo(&req)
	if err != nil {
		panic(err)
	}
	log.Print(vs)
}
