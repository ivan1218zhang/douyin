package dal

import "douyin/cmd/video/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
