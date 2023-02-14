package db

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

// MGetFollowedUser query list of FollowedUser
func MGetFollowedUser(ctx context.Context, userID int64) ([]int64, error) {
	var res []int64

	err := DB.WithContext(ctx).Model(&repository.FollowedUser{}).Where("user_id = ?", userID).Pluck("followed_user_id", &res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// MGetFollowerUser query list of FollowedUser
func MGetFollowerUser(ctx context.Context, userID int64) ([]int64, error) {
	var res []int64

	err := DB.WithContext(ctx).Model(&repository.FollowerUser{}).Where("user_id = ?", userID).Pluck("follower_user_id", &res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// MGetUsers query list of User
func MGetUsers(ctx context.Context, userIDs []int64) ([]*common.User, error) {
	res := make([]*common.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Table("users").Error; err != nil {
		return nil, err
	}
	return res, nil
}
