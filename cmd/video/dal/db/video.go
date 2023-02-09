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
	tm := time.Unix(0, latestTime*int64(time.Millisecond))
	if err := db.DB.WithContext(ctx).Where("created_at < ?", tm.Format("2006-01-02 03:04:05")).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
