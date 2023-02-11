package repository

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    int64
	VideoId   int64
	Content   string
}
