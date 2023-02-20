package db

import (
	"context"
	"douyin/pkg/repository"
)

func AddRelation(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Table("relation").Create(&repository.Relation{
		UserId:     toUserId,
		FollowerId: userId,
	}).Error
}

func DeleteRelation(ctx context.Context, userId int64, toUserId int64) error {
	return DB.WithContext(ctx).Table("relation").Where("user_id = ? and follower_id = ?", toUserId, userId).Delete(&repository.Relation{}).Error
}

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

// CountFollow 获取用户关注数
func CountFollow(ctx context.Context, userId int64) (int32, error) {
	var res int64
	err := DB.WithContext(ctx).Table("relation").Where("follower_id = ?", userId).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return int32(res), err
}

// CountFollower 获取用户粉丝数
func CountFollower(ctx context.Context, userId int64) (int32, error) {
	var res int64
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return int32(res), err
}

// IsFollow 判断当前用户是否关注目标用户
func IsFollow(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	var res int64
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ? and follower_id = ?", userId, toUserId).Count(&res).Error
	if err != nil {
		return false, err
	}
	return res > 0, err
}

type countResult struct {
	Count int32
	Id    int64
}

// MCountFollow 获取用户关注数
func MCountFollow(ctx context.Context, userIdList []int64) ([]int32, error) {
	var models []*countResult
	res := make([]int32, len(userIdList))
	err := DB.WithContext(ctx).Debug().Table("relation").Select("follower_id as id,count(1) as count").
		Where("follower_id in ?", userIdList).Group("follower_id").Scan(&models).Error
	if err != nil {
		return res, err
	}
	resMap := map[int64]*countResult{}
	for i := 0; i < len(models); i++ {
		resMap[models[i].Id] = models[i]
	}
	for i := 0; i < len(res); i++ {
		if resMap[userIdList[i]] != nil {
			res[i] = resMap[userIdList[i]].Count
		} else {
			res[i] = 0
		}
	}
	return res, err
}

// MCountFollower 获取用户粉丝数
func MCountFollower(ctx context.Context, userIdList []int64) ([]int32, error) {
	res := make([]int32, len(userIdList))
	var models []*countResult
	err := DB.WithContext(ctx).Debug().Table("relation").Select("user_id as id,count(1) as count").
		Where("user_id in ?", userIdList).Group("user_id").Scan(&models).Error
	if err != nil {
		return res, err
	}
	resMap := map[int64]*countResult{}
	for i := 0; i < len(models); i++ {
		resMap[models[i].Id] = models[i]
	}
	for i := 0; i < len(res); i++ {
		if resMap[userIdList[i]] != nil {
			res[i] = resMap[userIdList[i]].Count
		} else {
			res[i] = 0
		}
	}
	return res, err
}

// MIsFollow 判断当前用户是否关注目标用户
func MIsFollow(ctx context.Context, userId int64, toUserIdList []int64) ([]bool, error) {
	res := make([]bool, len(toUserIdList))
	var models []*countResult
	err := DB.WithContext(ctx).Debug().Table("relation").Select("user_id as id,count(1) as count").
		Where("user_id in ? and follower_id = ?", toUserIdList, userId).Group("user_id").Scan(&models).Error
	if err != nil {
		return res, err
	}
	resMap := map[int64]*countResult{}
	for i := 0; i < len(models); i++ {
		resMap[models[i].Id] = models[i]
	}
	for i := 0; i < len(res); i++ {
		if resMap[toUserIdList[i]] != nil {
			res[i] = resMap[toUserIdList[i]].Count > 0
		} else {
			res[i] = false
		}
	}
	return res, err
}
