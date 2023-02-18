package rdb

import (
	"douyin/pkg/constants"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

// 声明一个全局的redisDb变量
var (
	redisClient   *redis.Client
	initRedisOnce sync.Once
)

func GetRDB() *redis.Client {
	initRedisOnce.Do(initRDB)
	return redisClient
}

func initRDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress, // redis地址
		Password: "123456",               // redis密码，没有则留空
		DB:       7,                      // 默认数据库，默认是0
	})

	//redisClient.FlushDB()

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	log.Println("redis connected")
}
