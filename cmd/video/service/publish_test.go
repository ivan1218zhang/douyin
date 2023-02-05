package service

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/pkg/conf"
	"douyin/pkg/db"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestPublishService_Publish(t *testing.T) {
	// 创建视频保存的文件夹
	err := os.MkdirAll(conf.CDN.LocalPath, os.ModePerm)
	if err != nil {
		panic(err)
		return
	}
	// 初始化数据库
	err = db.Init()
	if err != nil {
		panic(err)
		return
	}
	// 读取测试视频
	file, err := os.Open("D:/test.mp4")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}
	// 生成请求
	req := &video.PublishReq{
		UserId: 0,
		Title:  "",
		Data:   data,
	}
	// 请求
	ctx := context.Background()
	publishService := NewPublishService(ctx)
	err = publishService.Publish(req)
	// 等待goroutine结束
	time.Sleep(10000 * time.Millisecond)
}
