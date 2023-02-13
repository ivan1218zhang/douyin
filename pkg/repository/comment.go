package repository

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    int64
	VideoId   int64
	Content   string
}
