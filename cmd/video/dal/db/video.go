package db

import (
	"context"
	"douyin/kitex_gen/common"
	"douyin/pkg/db"
	"douyin/pkg/repository"
	"gorm.io/gorm/clause"
	"time"
)

// QueryVideo query list of video info
func QueryVideo(ctx context.Context, userID int64) ([]*repository.Video, error) {
	var res []*repository.Video
	conn := db.DB.WithContext(ctx).Model(&repository.Video{}).Where("author_id = ?", userID)
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
	if err := db.DB.WithContext(ctx).Where("created_at < ?", tm.Format("2006-01-02 15:04:05")).Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
