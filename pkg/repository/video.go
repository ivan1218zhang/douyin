package repository

import (
	"douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Title         string
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}
