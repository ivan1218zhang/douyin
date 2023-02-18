package redis_test

import (
	"douyin/cmd/comment/dal/redis"
	"testing"
)

func TestInit(t *testing.T) {
	redis.Init()
}
