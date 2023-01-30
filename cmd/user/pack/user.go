package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{UserId: int64(u.ID), UserName: u.UserName}
}
