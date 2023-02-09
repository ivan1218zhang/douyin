package db

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/common"
	"douyin/pkg/repository"
)

// QueryVideo query list of video info
func QueryVideo(ctx context.Context, userID int64) ([]*common.Video, error) {
	var res []*common.Video
	conn := DB.WithContext(ctx).Model(&repository.Video{}).Where("user_id = ?", userID)
	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// CreateVideo create video
func CreateVideo(ctx context.Context, video *repository.Video) error {
	return db.DB.WithContext(ctx).Create(video).Error
}

// MGetVideo multiple get list of Video info
func MGetVideo(ctx context.Context, latestTime int64) ([]*common.Video, error) {
	var res []*common.Video

	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
