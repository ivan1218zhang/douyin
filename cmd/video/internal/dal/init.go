package dal

import (
	"douyin/cmd/video/internal/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
