package db

import (
	"context"
	"douyin/pkg/db"
	"douyin/pkg/repository"
)

func CreateFavorite(ctx context.Context, userID int64, videoID int64) error {
	err := db.DB.WithContext(ctx).Create(&repository.Favorite{
		UserId:  userID,
		VideoId: videoID,
	}).Error

	return err
}

func DeleteFavorite(ctx context.Context, userID int64, videoID int64) error {
	result := db.DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Delete(&repository.Favorite{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IsFavorite(ctx context.Context, userID int64, videoID int64) (bool, error) {
	var count int64
	err := db.DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Count(&count).Error
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

/*
	func MIsFavorite(ctx context.Context, userID int64, videoID []int64) (bool, error) {
		var count int64
		err := db.DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Count(&count).Error
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
	if err := db.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&fs).Error; err != nil {
		return nil, err
	}

	return fs, nil

}
