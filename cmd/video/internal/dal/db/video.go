package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorId      int64
	Author        User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}

// MGetVideos multiple get list of Video info
func MGetVideos(ctx context.Context, latestTime *int64) ([]*Video, error) {
	var res []*Video

	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
