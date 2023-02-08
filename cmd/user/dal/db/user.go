package db

import (
	"context"
	"douyin/pkg/repository"
)

func GetUsers(ctx context.Context, userID int64) (*repository.User, error) {
	res := &repository.User{}

	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *repository.User) (int64, error) {
	// 使用雪花ID生成器生成id
	user.ID = snowflake.NextSnowID()
	return user.ID, DB.WithContext(ctx).Create(user).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*repository.User, error) {
	res := make([]*repository.User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
