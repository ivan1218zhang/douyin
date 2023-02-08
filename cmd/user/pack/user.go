package pack

import (
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

// User pack user info
func User(u *repository.User) *common.User {
	if u == nil {
		return nil
	}
	return &common.User{Id: int64(u.ID), Name: u.UserName}
}
