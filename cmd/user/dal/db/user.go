package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string
	Password      string
	FollowCount   int64
	FollowerCount int64
}
