package db

import (
	"context"
	"douyin/pkg/repository"
)

func CreateFavorite(ctx context.Context, userID int64, videoID int64) error {
	err := DB.WithContext(ctx).Create(&repository.Favorite{
		UserId:  userID,
		VideoId: videoID,
	}).Error

	return err
}

func DeleteFavorite(ctx context.Context, userID int64, videoID int64) error {
	result := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Delete(&repository.Favorite{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func MIsFavorite(ctx context.Context, userID int64, videoIdList []int64) ([]bool, error) {
	res := make([]bool, len(videoIdList))
	var models []*countResult
	err := DB.WithContext(ctx).Table("favorite").Where("user_id = ? and video_id in ?", userID, videoIdList).Group("video_id").Select("count(1) as count, video_id as id").Find(&models).Error
	if err != nil {
		//log error
		return nil, err
	}
	resMap := map[int64]*countResult{}
	for i := 0; i < len(models); i++ {
		resMap[models[i].Id] = models[i]
	}
	for i := 0; i < len(videoIdList); i++ {
		if resMap[videoIdList[i]] != nil {
			res[i] = resMap[videoIdList[i]].Count > 0
		} else {
			res[i] = false
		}
	}
	return res, nil
}

/*
	func MIsFavorite(ctx context.Context, userID int64, videoID []int64) (bool, error) {
		var count int64
		err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Count(&count).Error
		if err != nil {
			//log error
			return false, err
		}
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
*/
func MGetFavoriteByUserID(ctx context.Context, userID int64) ([]repository.Favorite, error) {
	var fs []repository.Favorite
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&fs).Error; err != nil {
		return nil, err
	}
	return fs, nil
}

type countResult struct {
	Id    int64
	Count int64
}

func MCountFavorite(ctx context.Context, idList []int64) ([]int64, error) {
	res := make([]int64, len(idList))
	var models []*countResult
	err := DB.WithContext(ctx).Table("favorite").Where("video_id in ?", idList).Select("count(1) as count, video_id as id").Group("video_id").Find(&models).Error
	if err != nil {
		return res, err
	}
	resMap := map[int64]*countResult{}
	for i := 0; i < len(models); i++ {
		resMap[models[i].Id] = models[i]
	}
	for i := 0; i < len(idList); i++ {
		if resMap[idList[i]] != nil {
			res[i] = resMap[idList[i]].Count
		} else {
			res[i] = 0
		}
	}
	return res, nil
}
