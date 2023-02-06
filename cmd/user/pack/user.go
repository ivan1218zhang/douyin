package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/common"
)

// User pack user info
func User(u *db.User) *common.User {
	if u == nil {
		return nil
	}
	return &common.User{Id: int64(u.ID), Name: u.UserName}
}
