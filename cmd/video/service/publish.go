package service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"douyin/pkg/repository"
	"douyin/pkg/util"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"strconv"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

func (p *PublishService) Publish(req *video.PublishReq) error {
	// 生成视频信息
	videoId := db.Snowflake.NextSnowID()
	videoName := fmt.Sprintf("%d.mp4", videoId)
	coverName := fmt.Sprintf("%d.png", videoId)
	playUrl := constants.CDN.Url + videoName
	coverUrl := constants.CDN.Url + coverName
	// 存入数据库
	videoModel := &repository.Video{
		ID:       videoId,
		AuthorId: req.UserId,
		Title:    req.Title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	// 存到cdn
	go saveVideoCdn(videoId, videoName, coverName, req.Data)
	return db.CreateVideo(p.ctx, videoModel)
}

// 把视频和封面存到七牛云
func saveVideoCdn(videoId int64, videoName string, coverName string, data []byte) {
	// 视频存本地
	err := ioutil.WriteFile(constants.CDN.LocalPath+videoName, data, 0644)
	if err != nil {
		panic(err)
	}
	// 截取封面
	err = util.SavePicture(videoName, coverName)
	if err != nil {
		panic(err)
	}
	sendRabbitmq(videoId)
}

// 发送到消息队列
func sendRabbitmq(videoId int64) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"publish", // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	body := strconv.FormatInt(videoId, 10)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	fmt.Println("Message sent to the queue!")
}
