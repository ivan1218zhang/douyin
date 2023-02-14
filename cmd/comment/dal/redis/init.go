package redis

import (
	"context"
	"douyin/pkg/constants"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is the Redis client used for interacting with Redis.
var RedisClient *redis.Client

func Init() {
	// Create a new Redis client.
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress,
		Password: "",
		DB:       0,
	})

	// Test the Redis client connection.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}
