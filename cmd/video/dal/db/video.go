package db

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorId      int64
	Author        db.User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount string
	CommentCount  string
	IsFavorite    bool
	Title         string
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

// QueryVideo query list of video info
func QueryVideo(ctx context.Context, userID int64) ([]*Video, error) {
	var res []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("user_id = ?", userID)

	if err := conn.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
