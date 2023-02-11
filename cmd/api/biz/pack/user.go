package pack

import (
	common2 "douyin/cmd/api/biz/model/common"
	"douyin/kitex_gen/common"
)

func User(user *common.User) *common2.User {
	return &common2.User{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
