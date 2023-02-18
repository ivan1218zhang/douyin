package dal

import (
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/redis"
)

// Init init dal
func Init() {
	db.Init()    // mysql init
	redis.Init() // redis init
}
