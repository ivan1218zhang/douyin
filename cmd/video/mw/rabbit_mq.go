package mw

import (
	"douyin/pkg/util"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func InitRabbitmq() error {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//视频存七牛云
			err = util.UploadCdn(fmt.Sprintf("%s.mp4", string(d.Body)))
			if err != nil {
				panic(err)
			}
			//封面存七牛云
			err = util.UploadCdn(fmt.Sprintf("%s.png", string(d.Body)))
			if err != nil {
				panic(err)
			}
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
	return nil
}
