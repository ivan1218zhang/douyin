package db

import (
	"context"
	"douyin/pkg/repository"
)

func GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*repository.Comment, error) {
	res := make([]*repository.Comment, 0)

	if err := DB.WithContext(ctx).Where("videoid = ?", videoId).Find(&res).Error; err != nil {
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
