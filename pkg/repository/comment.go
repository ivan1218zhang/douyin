package repository

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserId    int64
	VideoId   int64  `gorm:"Index:idx_comment_videoid"`
	Content   string `gorm:"type:longtext"`
}
