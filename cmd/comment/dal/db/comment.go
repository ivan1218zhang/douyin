package db

import (
	"context"
	"douyin/pkg/repository"
)

func GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*repository.Comment, error) {
	res := make([]*repository.Comment, 0)

	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CreateComment(ctx context.Context, comment *repository.Comment) (*repository.Comment, error) {
	// 使用雪花ID生成器生成id
	comment.ID = snowflake.NextSnowID()
	return comment, DB.WithContext(ctx).Create(comment).Error
}

func DeleteComment(ctx context.Context, commentId int64) error {
	return DB.WithContext(ctx).Delete(&repository.Comment{ID: commentId}).Error
}

func GetCommentCountByVideoID(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&repository.Comment{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
