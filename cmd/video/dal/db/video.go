package db

import (
	"context"
	"douyin/pkg/db"
	"douyin/pkg/repository"
)

func CreateVideo(ctx context.Context, video *repository.Video) error {
	return db.DB.WithContext(ctx).Create(video).Error
}
