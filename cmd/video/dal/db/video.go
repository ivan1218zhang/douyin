package db

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/pkg/db"
	"douyin/pkg/repository"
	"time"
)

// QueryVideo query list of video info
func QueryVideo(ctx context.Context, userID int64) ([]*common.Video, error) {
	var res []*common.Video
	conn := db.DB.WithContext(ctx).Model(&repository.Video{}).Where("user_id = ?", userID)
	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// CreateVideo create video
func CreateVideo(ctx context.Context, video *repository.Video) error {
	return db.DB.WithContext(ctx).Create(video).Error
}

// GetVideo get video by id
func GetVideo(ctx context.Context, id int64) (repository.Video, error) {
	var v repository.Video
	err := db.DB.WithContext(ctx).First(&v, id).Error
	return v, err
}

// MGetVideo multiple get list of Video info
func MGetVideo(ctx context.Context, latestTime int64) (vs []*repository.Video, err error) {
	vs = []*repository.Video{}
	if latestTime != 0 {
		err = db.DB.WithContext(ctx).Where("created_at < ?", time.Unix(latestTime, 0)).Limit(10).Find(&vs).Error
	} else {
		err = db.DB.WithContext(ctx).Limit(10).Find(&vs).Error
	}

	return
}
